package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	"wafconsole/app/mqconsume/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData)

// Data .
type Data struct {
	log *log.Helper
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	l := log.NewHelper(log.With(logger, "module", "data"))
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		log: l,
	}, cleanup, nil
}
