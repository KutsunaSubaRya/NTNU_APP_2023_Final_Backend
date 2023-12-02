package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type GetCurriculumsResData struct {
	Number   string `json:"number"`
	Code     string `json:"code"`
	Name     string `json:"name"`
	Teacher  string `json:"teacher"`
	Time     string `json:"time"`
	Location string `json:"location"`
}

type GetCurriculumsReq struct {
	g.Meta `path:"/curriculums" tags:"CreateCurriculum" method:"GET"`
}

type GetCurriculumsRes struct {
	Error string                  `json:"error"`
	Data  []GetCurriculumsResData `json:"data"`
}
