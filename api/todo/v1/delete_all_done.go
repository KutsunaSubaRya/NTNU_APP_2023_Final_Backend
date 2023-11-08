package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type DeleteAllDoneReq struct {
	g.Meta `path:"/delete_all_done" tags:"DeleteAllDone" method:"GET"`
}

type DeleteAllDoneRes struct {
	Error string `json:"error"`
	Data  string `json:"data"`
}
