package v1

import "github.com/gogf/gf/v2/frame/g"

type GetNoteContentReq struct {
	g.Meta `path:"/note/content/{nid}" tags:"GetNoteContent" method:"GET"`
}

type GetNoteContentRes struct {
	Error string `json:"error"`
	Data  string `json:"data"`
}
