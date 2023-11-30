package note

import (
	"NTNU_APP_2023_Final_Backend/internal/dao"
	"NTNU_APP_2023_Final_Backend/internal/model/entity"
	"context"
	"github.com/gogf/gf/v2/frame/g"

	"NTNU_APP_2023_Final_Backend/api/note/v1"
)

func (c *ControllerV1) GetNoteContent(ctx context.Context, req *v1.GetNoteContentReq) (res *v1.GetNoteContentRes, err error) {
	r := g.RequestFromCtx(ctx)
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(v1.GetNoteContentRes{
			Error: err.Error(),
		})
	}

	uid := r.GetCtxVar("id").Int()
	nid := r.Get("nid").Int()

	note := entity.Notes{}
	m := dao.Notes.Ctx(ctx)
	err = m.Where("id", nid).Scan(&note)
	if err != nil {
		r.Response.WriteJsonExit(v1.GetNoteContentRes{
			Error: err.Error(),
		})
	}

	if note.UserId != uid {
		r.Response.WriteJsonExit(v1.GetNoteContentRes{
			Error: "permission denied",
		})
	}

	r.Response.WriteJson(v1.GetNoteContentRes{
		Data: note.Content,
	})
	return
}
