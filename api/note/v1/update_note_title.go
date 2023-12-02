package v1

import "github.com/gogf/gf/v2/frame/g"

type UpdateNoteTitleReq struct {
	g.Meta `path:"/note/title" tags:"UpdateNoteTitle" method:"PUT"`
	NoteId int    `p:"note_id" v:"required"`
	Title  string `p:"title" v:"required"`
}

type UpdateNoteTitleRes struct {
	Error string `json:"error"`
	Data  string `json:"data"`
}
