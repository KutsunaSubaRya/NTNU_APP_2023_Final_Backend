package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type SwitchTodoItemStatusReq struct {
	g.Meta `path:"/switch_todo_item_status" tags:"SwitchTodoItemStatus" method:"POST"`
	Id     int `p:"id" v:"required"`
}

type SwitchTodoItemStatusRes struct {
	Error string `json:"error"`
	Data  string `json:"data"`
}
