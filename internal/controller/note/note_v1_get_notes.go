package note

import (
	"NTNU_APP_2023_Final_Backend/internal/model/entity"
	"context"
	"github.com/gogf/gf/v2/frame/g"

	"NTNU_APP_2023_Final_Backend/api/note/v1"
)

func (c *ControllerV1) GetNotes(ctx context.Context, req *v1.GetNotesReq) (res *v1.GetNotesRes, err error) {
	r := g.RequestFromCtx(ctx)
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(v1.GetNotesRes{
			Error: err.Error(),
		})
	}

	uid := r.GetCtxVar("id").Int()
	var notes []entity.Notes
	err = g.Model("notes").Where("user_id", uid).Scan(&notes)
	if err != nil {
		r.Response.WriteJsonExit(v1.GetNotesRes{
			Error: err.Error(),
		})
	}

	var resData []v1.GetNotesResData
	for _, note := range notes {
		tmpData := v1.GetNotesResData{}
		tmpData.NoteId = note.Id
		tmpData.Title = note.Title
		resData = append(resData, tmpData)
	}

	r.Response.WriteJson(v1.GetNotesRes{
		Data: resData,
	})
	return
}
