package user

import (
	"net/http"
	"storm-server/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"storm-server/app/sys/cmd/api/internal/logic/sys/user"
	"storm-server/app/sys/cmd/api/internal/svc"
	"storm-server/app/sys/cmd/api/internal/types"
)

func UpdateUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateUserReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewUpdateUserLogic(r.Context(), svcCtx)
		resp, err := l.UpdateUser(&req)
		result.HttpResult(r, w, resp, err)
	}
}
