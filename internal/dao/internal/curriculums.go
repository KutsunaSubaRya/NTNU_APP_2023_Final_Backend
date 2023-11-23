// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CurriculumsDao is the data access object for table curriculums.
type CurriculumsDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns CurriculumsColumns // columns contains all the column names of Table for convenient usage.
}

// CurriculumsColumns defines and stores column names for table curriculums.
type CurriculumsColumns struct {
	Id        string //
	UserId    string //
	CourseId  string //
	CreatedAt string //
	UpdatedAt string //
	DeletedAt string //
}

// curriculumsColumns holds the columns for table curriculums.
var curriculumsColumns = CurriculumsColumns{
	Id:        "id",
	UserId:    "user_id",
	CourseId:  "course_id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewCurriculumsDao creates and returns a new DAO object for table data access.
func NewCurriculumsDao() *CurriculumsDao {
	return &CurriculumsDao{
		group:   "default",
		table:   "curriculums",
		columns: curriculumsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CurriculumsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CurriculumsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CurriculumsDao) Columns() CurriculumsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CurriculumsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CurriculumsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CurriculumsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
