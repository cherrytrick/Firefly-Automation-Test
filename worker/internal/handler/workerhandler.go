package handler

import (
	"net/http"

	"worker/internal/logic"
	"worker/internal/svc"
	"worker/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func WorkerHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewWorkerLogic(r.Context(), ctx)
		resp, err := l.Worker(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
