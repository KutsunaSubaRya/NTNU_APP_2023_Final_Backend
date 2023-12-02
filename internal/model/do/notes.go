// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Notes is the golang structure of table notes for DAO operations like Where/Data.
type Notes struct {
	g.Meta    `orm:"table:notes, do:true"`
	Id        interface{} //
	UserId    interface{} //
	Title     interface{} //
	Content   interface{} //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
	DeletedAt *gtime.Time //
}
