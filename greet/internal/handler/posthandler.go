package handler

import (
	"net/http"

	"github.com/smiecj/go_zero_demo/greet/internal/logic"
	"github.com/smiecj/go_zero_demo/greet/internal/svc"
	"github.com/smiecj/go_zero_demo/greet/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func PostHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Query
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewPostLogic(r.Context(), svcCtx)
		resp, err := l.Post(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
