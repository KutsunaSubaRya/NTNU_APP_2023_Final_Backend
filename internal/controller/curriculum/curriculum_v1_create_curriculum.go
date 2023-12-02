package curriculum

import (
	"NTNU_APP_2023_Final_Backend/internal/dao"
	"context"
	"github.com/gogf/gf/v2/frame/g"

	"NTNU_APP_2023_Final_Backend/api/curriculum/v1"
)

func (c *ControllerV1) CreateCurriculum(ctx context.Context, req *v1.CreateCurriculumReq) (res *v1.CreateCurriculumRes, err error) {
	r := g.RequestFromCtx(ctx)
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(v1.CreateCurriculumRes{
			Error: err.Error(),
		})
	}

	uid := r.GetCtxVar("id").Int()
	cid, _ := g.Model("courses").
		Fields("id").
		Where("number", req.CourseId).Value()

	m := dao.Curriculums.Ctx(ctx)
	cnt, err := m.Where("user_id=? AND course_id=?", uid, cid).Count()
	if cnt > 0 {
		r.Response.WriteJsonExit(v1.CreateCurriculumRes{
			Error: "Already exists",
		})
	}

	_, err = m.Data(g.Map{
		"user_id":   uid,
		"course_id": cid,
	}).Insert()
	if err != nil {
		r.Response.WriteJsonExit(v1.CreateCurriculumRes{
			Error: err.Error(),
		})
	}

	r.Response.WriteJson(v1.CreateCurriculumRes{
		Data: "Success",
	})
	return
}
