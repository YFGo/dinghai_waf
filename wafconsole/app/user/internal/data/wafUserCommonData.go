package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"

	"wafconsole/app/user/internal/biz"
	"wafconsole/utils/sendinfo"
)

type wafUserCommonRepo struct {
	data *Data
	log  *log.Helper
}

func NewWafUserCommonRepo(data *Data, logger log.Logger) biz.WafUserCommonRepo {
	return &wafUserCommonRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// SaveKVToRs 保存kv字符串到reids中
func (w *wafUserCommonRepo) SaveKVToRs(ctx context.Context, userEmail, systemCode string, expiration time.Duration) error {
	return w.data.rdb.Set(ctx, userEmail, systemCode, expiration).Err()
}

func (w *wafUserCommonRepo) SendSingEmailCode(ctx context.Context, content, userEmail string) error {
	emailCfg := w.data.emailCfg
	if err := sendinfo.SendEmail(emailCfg.SmtpServer, emailCfg.SmtpProt, emailCfg.SmtpUserName, userEmail, content, emailCfg.Auth); err != nil {
		w.log.WithContext(ctx).Error("send sing code is failed")
		return err
	}
	return nil
}
