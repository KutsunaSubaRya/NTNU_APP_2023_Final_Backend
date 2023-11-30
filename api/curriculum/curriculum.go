// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package curriculum

import (
	"context"

	"NTNU_APP_2023_Final_Backend/api/curriculum/v1"
)

type ICurriculumV1 interface {
	CreateCurriculum(ctx context.Context, req *v1.CreateCurriculumReq) (res *v1.CreateCurriculumRes, err error)
	DeleteCurriculum(ctx context.Context, req *v1.DeleteCurriculumReq) (res *v1.DeleteCurriculumRes, err error)
	GetCurriculums(ctx context.Context, req *v1.GetCurriculumsReq) (res *v1.GetCurriculumsRes, err error)
}
