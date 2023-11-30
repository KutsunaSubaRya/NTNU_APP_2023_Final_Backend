package v1

import "github.com/gogf/gf/v2/frame/g"

type CreateNoteReq struct {
	g.Meta `path:"/note" tags:"CreateNote" method:"POST"`
	Title  string `p:"title" v:"required"`
}

type CreateNoteRes struct {
	Error string `json:"error"`
	Data  string `json:"data"`
}
