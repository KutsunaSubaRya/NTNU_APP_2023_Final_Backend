package v1

import "github.com/gogf/gf/v2/frame/g"

type GetNotesResData struct {
	NoteId int    `json:"note_id"`
	Title  string `json:"title"`
}

type GetNotesReq struct {
	g.Meta `path:"/notes" tags:"GetNotes" method:"GET"`
}

type GetNotesRes struct {
	Error string            `json:"error"`
	Data  []GetNotesResData `json:"data"`
}
