package mycontext

import (
	"context"
	"encoding/base64"
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"

	"log/slog"
)

// CustomContext 自定义上下文类型
type CustomContext struct {
	// 内嵌标准Context
	context.Context

	// 元数据存储（线程安全）
	metadata  map[string]interface{}
	metaMutex sync.RWMutex

	// 链路追踪
	traceID    string
	spanID     string
	parentSpan string

	// 日志记录（支持结构化日志）
	logger *slog.Logger

	// 控制参数
	timeout    time.Duration
	deadline   *time.Time
	cancelFunc context.CancelFunc

	// 错误处理
	err      error
	errMutex sync.Mutex

	// 性能监控
	metricsCollector MetricsCollector
}

// MetricsCollector 指标收集接口
type MetricsCollector interface {
	RecordDuration(key string, dt time.Duration)
	IncrementCounter(key string)
}

type defaultMetricsCollector struct{}

func (d defaultMetricsCollector) RecordDuration(key string, dt time.Duration) {}
func (d defaultMetricsCollector) IncrementCounter(key string)                 {}

// NewAppCtx 构造函数
func NewAppCtx(ctx context.Context, opts ...Option) *CustomContext {
	baseCtx, cancel := context.WithCancel(ctx)
	cc := &CustomContext{
		Context:          baseCtx,
		metadata:         make(map[string]interface{}),
		logger:           defaultLogger(), // 默认使用 slog 的 JSONHandler
		cancelFunc:       cancel,
		metricsCollector: defaultMetricsCollector{},
	}

	for _, opt := range opts {
		opt(cc)
	}

	// 自动注入追踪ID（如果未设置）
	if cc.traceID == "" {
		cc.traceID = generateTraceID()
	}

	return cc
}

func defaultLogger() *slog.Logger {
	// os.Stout 日志打印在控制台上
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo, // 默认日志级别
	})
	return slog.New(handler)
}

// Option 模式配置函数
type Option func(*CustomContext)

func WithLogger(logger *slog.Logger) Option {
	return func(cc *CustomContext) {
		cc.logger = logger
	}
}

func WithTraceID(id string) Option {
	return func(cc *CustomContext) {
		cc.traceID = id
	}
}

// Deadline 实现标准Context接口
func (cc *CustomContext) Deadline() (time.Time, bool) {
	if cc.deadline != nil {
		return *cc.deadline, true
	}
	return cc.Context.Deadline()
}

func (cc *CustomContext) Done() <-chan struct{} {
	return cc.Context.Done()
}

func (cc *CustomContext) Err() error {
	cc.errMutex.Lock()
	defer cc.errMutex.Unlock()

	if cc.err != nil {
		return cc.err
	}
	return cc.Context.Err()
}

// Value 查找顺序：自定义元数据 -> 父Context
func (cc *CustomContext) Value(key interface{}) interface{} {
	cc.metaMutex.RLock()
	defer cc.metaMutex.RUnlock()

	if k, ok := key.(string); ok {
		if v, exists := cc.metadata[k]; exists {
			return v
		}
	}
	return cc.Context.Value(key)
}

/**************** 扩展功能 ****************/

// Set 元数据操作
func (cc *CustomContext) Set(key string, value interface{}) {
	cc.metaMutex.Lock()
	defer cc.metaMutex.Unlock()
	cc.metadata[key] = value
}

func (cc *CustomContext) Get(key string) (interface{}, bool) {
	cc.metaMutex.RLock()
	defer cc.metaMutex.RUnlock()
	v, exists := cc.metadata[key]
	return v, exists
}

// TraceID 链路追踪
func (cc *CustomContext) TraceID() string {
	return cc.traceID
}

func (cc *CustomContext) NewChildSpan() *CustomContext {
	newSpanID := generateSpanID()
	childLogger := cc.logger.With(
		slog.Group("span", slog.String("id", newSpanID)),
	)

	return &CustomContext{
		Context:    cc.Context,
		traceID:    cc.traceID,
		parentSpan: cc.spanID,
		spanID:     newSpanID,
		logger:     childLogger,
	}
}

// SetError 错误处理
func (cc *CustomContext) SetError(err error) {
	cc.errMutex.Lock()
	defer cc.errMutex.Unlock()

	if cc.err == nil { // 只记录第一个错误
		cc.err = err
		cc.logger.Error("context error",
			slog.String("traceID", cc.traceID),
			slog.String("error", err.Error()))

		// 上报错误指标
		cc.metricsCollector.IncrementCounter("context.errors")
	}
}

// Timeout 超时控制
func (cc *CustomContext) Timeout() time.Duration {
	return cc.timeout
}

func (cc *CustomContext) SetTimeout(d time.Duration) {
	cc.timeout = d
	if cc.deadline == nil {
		cc.deadline = new(time.Time)
	}
	*cc.deadline = time.Now().Add(d).UTC()

	// 自动超时取消
	go func() {
		select {
		case <-time.After(d):
			cc.SetError(fmt.Errorf("context timeout after %s", d))
			cc.cancelFunc()
		case <-cc.Done():
		}
	}()
}

// Log 日志增强
func (cc *CustomContext) Log() *slog.Logger {
	return cc.logger.With(
		slog.String("traceID", cc.traceID),
		slog.String("spanID", cc.spanID),
	)
}

// CancelWithReason 优雅取消（带原因）
func (cc *CustomContext) CancelWithReason(reason string) {
	if reason == "" {
		cc.cancelFunc()
		return
	}

	// 修改上下文的err为reason
	cc.SetError(fmt.Errorf("context cancelled with reason: %s", reason))
	cc.cancelFunc()

	// 上报取消指标
	cc.metricsCollector.IncrementCounter("context.cancellations")
}

// Clone 深度拷贝（用于并发安全传递）
func (cc *CustomContext) Clone() *CustomContext {
	cc.metaMutex.RLock()
	defer cc.metaMutex.RUnlock()

	newCc := &CustomContext{
		Context:          cc.Context,
		traceID:          cc.traceID,
		spanID:           cc.spanID,
		parentSpan:       cc.parentSpan,
		logger:           cc.logger,
		timeout:          cc.timeout,
		metricsCollector: cc.metricsCollector,
	}

	// 深拷贝元数据
	newCc.metadata = make(map[string]interface{}, len(cc.metadata))
	for k, v := range cc.metadata {
		newCc.metadata[k] = deepCopy(v)
	}

	return newCc
}

// 辅助函数
func generateTraceID() string {
	// 使用随机数生成分布式ID的一部分
	return fmt.Sprintf("trace-%d-%x", time.Now().UnixNano(), rand.Int63())
}

func generateSpanID() string {
	// 可以用随机数或者更复杂的算法生成
	return fmt.Sprintf("span-%x", rand.Int63())
}

func deepCopy(src interface{}) interface{} {
	// 实现简单的深度拷贝（这里仅支持基本类型和字符串）
	switch v := src.(type) {
	case string:
		return v
	case int, int8, int16, int32, int64:
		return v
	case uint, uint8, uint16, uint32, uint64:
		return v
	case float32, float64:
		return v
	case []byte:
		return base64.StdEncoding.EncodeToString(v)
	default:
		return src // 其他类型不处理
	}
}
