// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Todos is the golang structure for table todos.
type Todos struct {
	Id        int         `json:"id"        ` //
	UserId    int         `json:"userId"    ` //
	Content   string      `json:"content"   ` //
	IsDone    bool        `json:"isDone"    ` //
	CreatedAt *gtime.Time `json:"createdAt" ` //
	UpdatedAt *gtime.Time `json:"updatedAt" ` //
	DeletedAt *gtime.Time `json:"deletedAt" ` //
}
