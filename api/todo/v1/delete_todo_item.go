package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type DeleteTodoItemReq struct {
	g.Meta `path:"/delete_todo_item" tags:"DeleteTodoItem" method:"POST"`
	Id     int `p:"id" v:"required"`
}

type DeleteTodoItemRes struct {
	Error string `json:"error"`
	Data  string `json:"data"`
}
