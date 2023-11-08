// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package todo

import (
	"context"
	
	"NTNU_APP_2023_Final_Backend/api/todo/v1"
)

type ITodoV1 interface {
	CreateTodoItem(ctx context.Context, req *v1.CreateTodoItemReq) (res *v1.CreateTodoItemRes, err error)
	DeleteAllDone(ctx context.Context, req *v1.DeleteAllDoneReq) (res *v1.DeleteAllDoneRes, err error)
	DeleteTodoItem(ctx context.Context, req *v1.DeleteTodoItemReq) (res *v1.DeleteTodoItemRes, err error)
	SwitchTodoItemStatus(ctx context.Context, req *v1.SwitchTodoItemStatusReq) (res *v1.SwitchTodoItemStatusRes, err error)
}


