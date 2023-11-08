package todo

import (
	"NTNU_APP_2023_Final_Backend/internal/dao"
	"NTNU_APP_2023_Final_Backend/internal/model/entity"
	"context"
	"github.com/gogf/gf/v2/frame/g"

	"NTNU_APP_2023_Final_Backend/api/todo/v1"
)

func (c *ControllerV1) SwitchTodoItemStatus(ctx context.Context, req *v1.SwitchTodoItemStatusReq) (res *v1.SwitchTodoItemStatusRes, err error) {
	r := g.RequestFromCtx(ctx)
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(v1.SwitchTodoItemStatusRes{
			Error: err.Error(),
		})
	}

	var todoItem entity.Todos
	m := dao.Todos.Ctx(ctx)
	err = m.Where("id", req.Id).Scan(&todoItem)
	if err != nil {
		r.Response.WriteJsonExit(v1.SwitchTodoItemStatusRes{
			Error: err.Error(),
		})
	}
	
	if todoItem.UserId != r.GetCtxVar("id").Int() {
		r.Response.WriteJsonExit(v1.SwitchTodoItemStatusRes{
			Error: "permission denied",
		})
	}

	_, err = m.Update(
		g.Map{
			"is_done": !todoItem.IsDone,
		},
		"id", req.Id,
	)
	if err != nil {
		r.Response.WriteJsonExit(v1.SwitchTodoItemStatusRes{
			Error: err.Error(),
		})
	}

	r.Response.WriteJson(v1.SwitchTodoItemStatusRes{
		Data: "Success",
	})
	return
}
