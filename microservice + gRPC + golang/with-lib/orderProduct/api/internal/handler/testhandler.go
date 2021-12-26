package handler

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"orderProduct/api/internal/logic"
	"orderProduct/api/internal/svc"
	"orderProduct/api/internal/types"
)

func TestHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TestGetReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewTestLogic(r.Context(), ctx)
		resp, err := l.Test(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
