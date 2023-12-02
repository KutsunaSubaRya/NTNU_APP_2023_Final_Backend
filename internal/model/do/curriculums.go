// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Curriculums is the golang structure of table curriculums for DAO operations like Where/Data.
type Curriculums struct {
	g.Meta    `orm:"table:curriculums, do:true"`
	Id        interface{} //
	UserId    interface{} //
	CourseId  interface{} //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
	DeletedAt *gtime.Time //
}
