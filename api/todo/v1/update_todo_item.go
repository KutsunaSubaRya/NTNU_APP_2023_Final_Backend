package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type UpdateTodoItemReq struct {
	g.Meta  `path:"/update_todo_item" tags:"UpdateTodoItem" method:"POST"`
	Id      int    `p:"id" v:"required"`
	Content string `p:"content" v:"required|length:1,50"`
	IsDone  bool   `p:"is_done" v:"required|boolean"`
}

type UpdateTodoItemRes struct {
	Error string `json:"error"`
	Data  string `json:"data"`
}
