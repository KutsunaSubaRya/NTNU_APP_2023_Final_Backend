package todo

import (
	"NTNU_APP_2023_Final_Backend/internal/dao"
	"context"
	"github.com/gogf/gf/v2/frame/g"

	"NTNU_APP_2023_Final_Backend/api/todo/v1"
)

func (c *ControllerV1) GetTodoList(ctx context.Context, req *v1.GetTodoListReq) (res *v1.GetTodoListRes, err error) {
	r := g.RequestFromCtx(ctx)
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(v1.GetTodoListRes{
			Error: err.Error(),
		})
	}

	uid := r.GetCtxVar("id").Int()
	var resData []v1.GetTodoListResData
	m := dao.Todos.Ctx(ctx)
	err = m.Where("user_id", uid).Scan(&resData)
	if err != nil {
		r.Response.WriteJsonExit(v1.GetTodoListRes{
			Error: err.Error(),
		})
	}

	r.Response.WriteJson(v1.GetTodoListRes{
		Data: resData,
	})
	return
}
