package permission

import (
	"net/http"
	"storm-server/common/result"

	"storm-server/app/sys/cmd/api/internal/logic/sys/permission"
	"storm-server/app/sys/cmd/api/internal/svc"
)

func RouteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := permission.NewRouteLogic(r.Context(), svcCtx)
		resp, err := l.Route()
		result.HttpResult(r, w, resp, err)
	}
}
