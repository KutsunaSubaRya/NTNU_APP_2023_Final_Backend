package v1

import "github.com/gogf/gf/v2/frame/g"

type UpdateNoteContentReq struct {
	g.Meta  `path:"/note/content" tags:"UpdateNoteContent" method:"PUT"`
	NoteId  int    `p:"note_id" v:"required"`
	Content string `p:"content" v:"required"`
}

type UpdateNoteContentRes struct {
	Error string `json:"error"`
	Data  string `json:"data"`
}
