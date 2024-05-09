package department

import (
	"net/http"
	"storm-server/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"storm-server/app/sys/cmd/api/internal/logic/sys/department"
	"storm-server/app/sys/cmd/api/internal/svc"
	"storm-server/app/sys/cmd/api/internal/types"
)

func DeleteDepartmentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteDepartmentReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := department.NewDeleteDepartmentLogic(r.Context(), svcCtx)
		resp, err := l.DeleteDepartment(&req)
		result.HttpResult(r, w, resp, err)
	}
}
