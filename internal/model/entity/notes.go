// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Notes is the golang structure for table notes.
type Notes struct {
	Id        int         `json:"id"        ` //
	UserId    int         `json:"userId"    ` //
	Title     string      `json:"title"     ` //
	Content   string      `json:"content"   ` //
	CreatedAt *gtime.Time `json:"createdAt" ` //
	UpdatedAt *gtime.Time `json:"updatedAt" ` //
	DeletedAt *gtime.Time `json:"deletedAt" ` //
}
