package curriculum

import (
	"NTNU_APP_2023_Final_Backend/internal/dao"
	"context"
	"github.com/gogf/gf/v2/frame/g"

	"NTNU_APP_2023_Final_Backend/api/curriculum/v1"
)

func (c *ControllerV1) DeleteCurriculum(ctx context.Context, req *v1.DeleteCurriculumReq) (res *v1.DeleteCurriculumRes, err error) {
	r := g.RequestFromCtx(ctx)
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(v1.DeleteCurriculumRes{
			Error: err.Error(),
		})
	}

	uid := r.GetCtxVar("id").Int()
	cid, _ := g.Model("courses").
		Fields("id").
		Where("number", req.CourseId).Value()

	m := dao.Curriculums.Ctx(ctx)
	_, err = m.Where("user_id=? AND course_id=?", uid, cid).Delete()
	if err != nil {
		r.Response.WriteJsonExit(v1.DeleteCurriculumRes{
			Error: err.Error(),
		})
	}

	r.Response.WriteJson(v1.DeleteCurriculumRes{
		Data: "Success",
	})
	return
}
