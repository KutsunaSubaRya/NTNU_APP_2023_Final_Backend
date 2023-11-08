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
}


