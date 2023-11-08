package todo

import (
	"NTNU_APP_2023_Final_Backend/internal/dao"
	"context"
	"github.com/gogf/gf/v2/frame/g"

	"NTNU_APP_2023_Final_Backend/api/todo/v1"
)

func (c *ControllerV1) DeleteTodoItem(ctx context.Context, req *v1.DeleteTodoItemReq) (res *v1.DeleteTodoItemRes, err error) {
	r := g.RequestFromCtx(ctx)
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(v1.DeleteTodoItemRes{
			Error: err.Error(),
		})
	}

	m := dao.Todos.Ctx(ctx)
	uid, err := m.Fields("user_id").Where("id", req.Id).Value()
	if err != nil {
		r.Response.WriteJsonExit(v1.DeleteTodoItemRes{
			Error: err.Error(),
		})
	}
	
	if uid.Int() != r.GetCtxVar("id").Int() {
		r.Response.WriteJsonExit(v1.DeleteTodoItemRes{
			Error: "permission denied",
		})
	}

	_, err = m.Where("id", req.Id).Delete()
	if err != nil {
		r.Response.WriteJsonExit(v1.DeleteTodoItemRes{
			Error: err.Error(),
		})
	}

	r.Response.WriteJson(v1.DeleteTodoItemRes{
		Data: "Success",
	})
	return
}
