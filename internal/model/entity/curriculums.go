// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Curriculums is the golang structure for table curriculums.
type Curriculums struct {
	Id        int         `json:"id"        ` //
	UserId    int         `json:"userId"    ` //
	CourseId  int         `json:"courseId"  ` //
	CreatedAt *gtime.Time `json:"createdAt" ` //
	UpdatedAt *gtime.Time `json:"updatedAt" ` //
	DeletedAt *gtime.Time `json:"deletedAt" ` //
}
