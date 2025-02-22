package utils

import (
	"encoding/json"
	stdhttp "net/http"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/transport/http"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Reply struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func ResponseEncoder(w http.ResponseWriter, r *http.Request, v any) error {
	data, err := json.Marshal(&Reply{
		Code: 0,
		Msg:  "",
		Data: v,
	})
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
	return nil
}

func ErrorEncoder(w http.ResponseWriter, r *http.Request, err error) {
	e := errors.FromError(err)
	code := int(e.Code)
	rep := &Reply{
		Code: e.Code,
		Msg:  e.Message,
		Data: e.Metadata,
	}

	body, err := json.Marshal(rep)
	if err != nil {
		w.WriteHeader(stdhttp.StatusInternalServerError)
		return
	}
	//参数错误 , 单独处理
	if codes.InvalidArgument == status.Code(e) {
		w.WriteHeader(stdhttp.StatusOK)
		w.Write(body)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if 0 < code && code <= 600 {
		w.WriteHeader(code)
	} else {
		w.WriteHeader(stdhttp.StatusOK)
	}
	w.Write(body)
}
