package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type CreateTodoItemReq struct {
	g.Meta  `path:"/create_todo_item" tags:"CreateTodoItem" method:"POST"`
	Content string `p:"content" v:"required|length:1,50"`
}

type CreateTodoItemRes struct {
	Error string `json:"error"`
	Data  string `json:"data"`
}
