package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type DeleteCurriculumReq struct {
	g.Meta   `path:"/curriculum" tags:"DeleteCurriculum" method:"DELETE"`
	CourseId string `p:"course_id" v:"required"`
}

type DeleteCurriculumRes struct {
	Error string `json:"error"`
	Data  string `json:"data"`
}
