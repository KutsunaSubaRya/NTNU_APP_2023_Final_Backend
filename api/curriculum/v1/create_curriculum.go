package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type CreateCurriculumReq struct {
	g.Meta   `path:"/curriculum" tags:"CreateCurriculum" method:"POST"`
	CourseId string `p:"course_id" v:"required"`
}

type CreateCurriculumRes struct {
	Error string `json:"error"`
	Data  string `json:"data"`
}
