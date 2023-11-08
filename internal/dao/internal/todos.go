// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TodosDao is the data access object for table todos.
type TodosDao struct {
	table   string       // table is the underlying table name of the DAO.
	group   string       // group is the database configuration group name of current DAO.
	columns TodosColumns // columns contains all the column names of Table for convenient usage.
}

// TodosColumns defines and stores column names for table todos.
type TodosColumns struct {
	Id      string //
	UserId  string //
	Content string //
	IsDone  string //
}

// todosColumns holds the columns for table todos.
var todosColumns = TodosColumns{
	Id:      "id",
	UserId:  "user_id",
	Content: "content",
	IsDone:  "is_done",
}

// NewTodosDao creates and returns a new DAO object for table data access.
func NewTodosDao() *TodosDao {
	return &TodosDao{
		group:   "default",
		table:   "todos",
		columns: todosColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TodosDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TodosDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TodosDao) Columns() TodosColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TodosDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TodosDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TodosDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
