package note

import (
	"NTNU_APP_2023_Final_Backend/internal/dao"
	"context"
	"github.com/gogf/gf/v2/frame/g"

	"NTNU_APP_2023_Final_Backend/api/note/v1"
)

func (c *ControllerV1) CreateNote(ctx context.Context, req *v1.CreateNoteReq) (res *v1.CreateNoteRes, err error) {
	r := g.RequestFromCtx(ctx)
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(v1.CreateNoteRes{
			Error: err.Error(),
		})
	}

	m := dao.Notes.Ctx(ctx)
	_, err = m.Data(g.Map{
		"user_id": r.GetCtxVar("id").Int(),
		"title":   req.Title,
	}).Insert()
	if err != nil {
		r.Response.WriteJsonExit(v1.CreateNoteRes{
			Error: err.Error(),
		})
	}

	r.Response.WriteJson(v1.CreateNoteRes{
		Data: "Success",
	})
	return
}
