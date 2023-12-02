package note

import (
	"NTNU_APP_2023_Final_Backend/internal/dao"
	"context"
	"github.com/gogf/gf/v2/frame/g"

	"NTNU_APP_2023_Final_Backend/api/note/v1"
)

func (c *ControllerV1) UpdateNoteContent(ctx context.Context, req *v1.UpdateNoteContentReq) (res *v1.UpdateNoteContentRes, err error) {
	r := g.RequestFromCtx(ctx)
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(v1.UpdateNoteContentRes{
			Error: err.Error(),
		})
	}

	id := r.GetCtxVar("id").Int()

	m := dao.Notes.Ctx(ctx)
	uid, err := m.Fields("user_id").Where("id", req.NoteId).Value()
	if err != nil {
		r.Response.WriteJsonExit(v1.UpdateNoteContentRes{
			Error: err.Error(),
		})
	}
	if uid.Int() != id {
		r.Response.WriteJsonExit(v1.UpdateNoteContentRes{
			Error: "permission denied",
		})
	}

	_, err = m.Data(g.Map{
		"content": req.Content,
	}).Where("id", req.NoteId).Update()
	if err != nil {
		r.Response.WriteJsonExit(v1.UpdateNoteContentRes{
			Error: err.Error(),
		})
	}

	r.Response.WriteJson(v1.UpdateNoteContentRes{
		Data: "Success",
	})
	return
}
