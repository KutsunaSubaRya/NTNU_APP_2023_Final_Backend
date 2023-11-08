// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Todos is the golang structure of table todos for DAO operations like Where/Data.
type Todos struct {
	g.Meta    `orm:"table:todos, do:true"`
	Id        interface{} //
	UserId    interface{} //
	Content   interface{} //
	IsDone    interface{} //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
	DeletedAt *gtime.Time //
}
