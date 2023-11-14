package todo

import (
	"NTNU_APP_2023_Final_Backend/internal/dao"
	"context"
	"github.com/gogf/gf/v2/frame/g"

	"NTNU_APP_2023_Final_Backend/api/todo/v1"
)

func (c *ControllerV1) CreateTodoItem(ctx context.Context, req *v1.CreateTodoItemReq) (res *v1.CreateTodoItemRes, err error) {
	r := g.RequestFromCtx(ctx)
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(v1.CreateTodoItemRes{
			Error: err.Error(),
		})
	}

	m := dao.Todos.Ctx(ctx)
	_, err = m.Data(g.Map{
		"user_id": r.GetCtxVar("id").Int(),
		"content": req.Content,
	}).Insert()
	if err != nil {
		r.Response.WriteJsonExit(v1.CreateTodoItemRes{
			Error: err.Error(),
		})
	}

	r.Response.WriteJson(v1.CreateTodoItemRes{
		Data: "Success",
	})
	return
}
