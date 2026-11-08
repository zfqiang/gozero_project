package handler

import (
	"net/http"
	"video/internal/logic"
	"video/internal/svc"
	"video/internal/types"

	"response" // ①

	"github.com/zeromicro/go-zero/rest/httpx"
)

func getVideoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetVideoLogic(r.Context(), svcCtx)
		resp, err := l.GetVideo(&req)
		response.Response(w, resp, err) //②

	}
}
