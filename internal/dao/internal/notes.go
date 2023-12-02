// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// NotesDao is the data access object for table notes.
type NotesDao struct {
	table   string       // table is the underlying table name of the DAO.
	group   string       // group is the database configuration group name of current DAO.
	columns NotesColumns // columns contains all the column names of Table for convenient usage.
}

// NotesColumns defines and stores column names for table notes.
type NotesColumns struct {
	Id        string //
	UserId    string //
	Title     string //
	Content   string //
	CreatedAt string //
	UpdatedAt string //
	DeletedAt string //
}

// notesColumns holds the columns for table notes.
var notesColumns = NotesColumns{
	Id:        "id",
	UserId:    "user_id",
	Title:     "title",
	Content:   "content",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewNotesDao creates and returns a new DAO object for table data access.
func NewNotesDao() *NotesDao {
	return &NotesDao{
		group:   "default",
		table:   "notes",
		columns: notesColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *NotesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *NotesDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *NotesDao) Columns() NotesColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *NotesDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *NotesDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *NotesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
