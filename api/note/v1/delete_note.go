package v1

import "github.com/gogf/gf/v2/frame/g"

type DeleteNoteReq struct {
	g.Meta `path:"/note" tags:"DeleteNote" method:"DELETE"`
	NoteId int `p:"note_id" v:"required"`
}

type DeleteNoteRes struct {
	Error string `json:"error"`
	Data  string `json:"data"`
}
