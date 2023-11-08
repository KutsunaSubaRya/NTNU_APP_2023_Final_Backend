package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type GetTodoListResData struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	IsDone  bool   `json:"is_done"`
}

type GetTodoListReq struct {
	g.Meta `path:"/get_todo_list" tags:"GetTodoList" method:"GET"`
}

type GetTodoListRes struct {
	Error string               `json:"error"`
	Data  []GetTodoListResData `json:"data"`
}
