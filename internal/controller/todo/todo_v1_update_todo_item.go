package todo

import (
	"NTNU_APP_2023_Final_Backend/internal/dao"
	"context"
	"github.com/gogf/gf/v2/frame/g"

	"NTNU_APP_2023_Final_Backend/api/todo/v1"
)

func (c *ControllerV1) UpdateTodoItem(ctx context.Context, req *v1.UpdateTodoItemReq) (res *v1.UpdateTodoItemRes, err error) {
	r := g.RequestFromCtx(ctx)
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(v1.UpdateTodoItemRes{
			Error: err.Error(),
		})
	}

	m := dao.Todos.Ctx(ctx)
	uid, err := m.Fields("user_id").Where("id", req.Id).Value()
	if err != nil {
		r.Response.WriteJsonExit(v1.UpdateTodoItemRes{
			Error: err.Error(),
		})
	}

	if uid.Int() == 0 {
		r.Response.WriteJsonExit(v1.UpdateTodoItemRes{
			Error: "todo item not found",
		})
	}

	if uid.Int() != r.GetCtxVar("id").Int() {
		r.Response.WriteJsonExit(v1.UpdateTodoItemRes{
			Error: "permission denied",
		})
	}

	_, err = m.Update(
		g.Map{
			"content": req.Content,
			"is_done": req.IsDone,
		},
		"id", req.Id,
	)
	if err != nil {
		r.Response.WriteJsonExit(v1.UpdateTodoItemRes{
			Error: err.Error(),
		})
	}

	r.Response.WriteJson(v1.UpdateTodoItemRes{
		Data: "Success",
	})
	return
}
