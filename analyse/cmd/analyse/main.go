package main

import (
	"analyse/internal/biz/attack_log"
	"analyse/internal/conf"
	"analyse/internal/data"
	"flag"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	_ "go.uber.org/automaxprocs"
)

var (
	flagconf string
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func wireApp(confServer *conf.Server, bootstrap *conf.Bootstrap) (func(), *attack_log.AttackLogUsercase, error) {
	dataData, cleanup, err := data.NewData(confServer, bootstrap)
	if err != nil {
		return nil, nil, err
	}
	attackLogRepo := data.NewAttackLogRepo(dataData)
	attackLogUsercase := attack_log.NewAttackLogUsercase(attackLogRepo)
	return func() {
		cleanup()
	}, attackLogUsercase, nil
}

func main() {
	flag.Parse()
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	cleanup, attackLogEvent, err := wireApp(bc.Server, &bc)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	//消费kafka消息
	end := attackLogEvent.ConsumerAttackEvents()

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-sigterm:
		slog.Warn("terminating: via signal")
	}
	end()
}
