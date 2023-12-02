package note

import (
	"NTNU_APP_2023_Final_Backend/internal/dao"
	"NTNU_APP_2023_Final_Backend/internal/model/entity"
	"context"
	"github.com/gogf/gf/v2/frame/g"

	"NTNU_APP_2023_Final_Backend/api/note/v1"
)

func (c *ControllerV1) DeleteNote(ctx context.Context, req *v1.DeleteNoteReq) (res *v1.DeleteNoteRes, err error) {
	r := g.RequestFromCtx(ctx)
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(v1.DeleteNoteRes{
			Error: err.Error(),
		})
	}

	uid := r.GetCtxVar("id").Int()
	nid := req.NoteId

	m := dao.Notes.Ctx(ctx)

	note := entity.Notes{}
	err = m.Where("id", nid).Scan(&note)
	if err != nil {
		r.Response.WriteJsonExit(v1.DeleteNoteRes{
			Error: err.Error(),
		})
	}

	if note.UserId != uid {
		r.Response.WriteJsonExit(v1.DeleteNoteRes{
			Error: "permission denied",
		})
	}

	_, err = m.Where("id=? AND user_id=?", nid, uid).Delete()
	if err != nil {
		r.Response.WriteJsonExit(v1.DeleteNoteRes{
			Error: err.Error(),
		})
	}

	r.Response.WriteJson(v1.DeleteNoteRes{
		Data: "Success",
	})
	return
}
