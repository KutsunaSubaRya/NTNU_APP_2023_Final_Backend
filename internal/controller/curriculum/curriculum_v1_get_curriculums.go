package curriculum

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"

	"NTNU_APP_2023_Final_Backend/api/curriculum/v1"
)

func (c *ControllerV1) GetCurriculums(ctx context.Context, req *v1.GetCurriculumsReq) (res *v1.GetCurriculumsRes, err error) {
	r := g.RequestFromCtx(ctx)
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(v1.GetCurriculumsRes{
			Error: err.Error(),
		})
	}

	uid := r.GetCtxVar("id").Int()

	courseIds, err := g.Model("curriculums").
		Fields("course_id").Where("user_id", uid).
		Distinct().OrderAsc("course_id").All()
	if err != nil {
		r.Response.WriteJsonExit(v1.GetCurriculumsRes{
			Error: err.Error(),
		})
	}
	var resData []v1.GetCurriculumsResData
	for _, courseId := range courseIds {
		courseIdInt := courseId["course_id"].Int()
		tmpData := v1.GetCurriculumsResData{}
		err := g.Model("courses").Where("id", courseIdInt).Scan(&tmpData)
		if err != nil {
			r.Response.WriteJsonExit(v1.GetCurriculumsRes{
				Error: err.Error(),
			})
		}
		resData = append(resData, tmpData)
	}

	r.Response.WriteJson(v1.GetCurriculumsRes{
		Data: resData,
	})
	return
}
