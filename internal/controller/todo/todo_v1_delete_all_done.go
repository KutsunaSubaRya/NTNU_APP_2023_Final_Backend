package todo

import (
	"NTNU_APP_2023_Final_Backend/internal/dao"
	"context"
	"github.com/gogf/gf/v2/frame/g"

	"NTNU_APP_2023_Final_Backend/api/todo/v1"
)

func (c *ControllerV1) DeleteAllDone(ctx context.Context, req *v1.DeleteAllDoneReq) (res *v1.DeleteAllDoneRes, err error) {
	r := g.RequestFromCtx(ctx)
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(v1.DeleteAllDoneRes{
			Error: err.Error(),
		})
	}

	id := r.GetCtxVar("id").Int()
	m := dao.Todos.Ctx(ctx)

	_, err = m.Where("user_id", id).Where("is_done", 1).Delete()
	if err != nil {
		r.Response.WriteJsonExit(v1.DeleteAllDoneRes{
			Error: err.Error(),
		})
	}

	r.Response.WriteJson(v1.DeleteAllDoneRes{
		Data: "Success",
	})
	return
}
