package biz

import (
	"context"
	"encoding/json"
	"log/slog"

	"wafcoraza/data/model"
)

type NormalHttpRepo interface {
	// SaveNormalHttpToKafka 将正常的http请求保存到kafka中
	SaveNormalHttpToKafka(ctx context.Context, normalHttpInfoJson, id string) error
}

type NormalHttpUsercase struct {
	repo NormalHttpRepo
}

func NewNormalHttpUsercase(repo NormalHttpRepo) *NormalHttpUsercase {
	return &NormalHttpUsercase{repo: repo}
}

// SaveNormalHttp 保存正常的http请求
func (n *NormalHttpUsercase) SaveNormalHttp(ctx context.Context, normalHttpInfo model.ListNormalHttp) error {
	normalHttpInfoJson, err := json.Marshal(&normalHttpInfo)
	if err != nil {
		slog.ErrorContext(ctx, "json marshal normal_http_info is error", err)
		return err
	}
	return n.repo.SaveNormalHttpToKafka(ctx, string(normalHttpInfoJson), normalHttpInfo.ID)
}
