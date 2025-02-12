package utils

import (
	"math/rand"
)

// LoadBalance 接口定义
type LoadBalance interface {
	// Next 参数remove是需要排除选择的后端Server
	Next(remove []string) *Server
	// UpdateServers 更新可用Server列表
	UpdateServers(servers []*Server)
}

// Server 后端Server定义
type Server struct {
	//主机地址
	Host string
	//主机名
	Name string
	Id   int
	//主机是否在线
	Online bool
}
type LoadBalanceRandom struct {
	servers []*Server
}

// NewLoadBalanceRandom 实例化 随机均衡负载
func NewLoadBalanceRandom(servers []*Server) *LoadBalanceRandom {
	newBalance := &LoadBalanceRandom{}
	newBalance.UpdateServers(servers)
	return newBalance
}

// Next 选择一个后端Server
func (r *LoadBalanceRandom) Next() *Server {
	if len(r.servers) == 0 {
		return nil
	}
	curIndex := rand.Intn(len(r.servers))
	return r.servers[curIndex]
}
func (r *LoadBalanceRandom) Get(key string) (*Server, error) {
	return r.Next(), nil
}

// UpdateServers 系统运行过程中，后端可用Server会更新
func (this *LoadBalanceRandom) UpdateServers(servers []*Server) {
	newServers := make([]*Server, 0)
	for _, e := range servers {
		if e.Online == true {
			newServers = append(newServers, e)
		}
	}
	this.servers = newServers
}
