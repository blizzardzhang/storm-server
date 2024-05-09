package client

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"storm-server/app/sys/cmd/api/internal/logic/sys/client"
	"storm-server/app/sys/cmd/api/internal/svc"
	"storm-server/app/sys/cmd/api/internal/types"
)

func ListClientHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListClientReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := client.NewListClientLogic(r.Context(), svcCtx)
		resp, err := l.ListClient(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
