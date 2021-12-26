package handler

import (
	"fmt"
	"net/http"

	"orderProduct/api/internal/logic"
	"orderProduct/api/internal/svc"
	"orderProduct/api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func OrderHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OrderReq
		if err := httpx.Parse(r, &req); err != nil {
			fmt.Println("check", req)
			httpx.Error(w, err)
			return
		}

		l := logic.NewOrderLogic(r.Context(), ctx)
		resp, err := l.Order(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
