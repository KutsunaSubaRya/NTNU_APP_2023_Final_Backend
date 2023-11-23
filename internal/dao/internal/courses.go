// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CoursesDao is the data access object for table courses.
type CoursesDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns CoursesColumns // columns contains all the column names of Table for convenient usage.
}

// CoursesColumns defines and stores column names for table courses.
type CoursesColumns struct {
	Id       string //
	Number   string //
	Code     string //
	Name     string //
	Teacher  string //
	Time     string //
	Location string //
}

// coursesColumns holds the columns for table courses.
var coursesColumns = CoursesColumns{
	Id:       "id",
	Number:   "number",
	Code:     "code",
	Name:     "name",
	Teacher:  "teacher",
	Time:     "time",
	Location: "location",
}

// NewCoursesDao creates and returns a new DAO object for table data access.
func NewCoursesDao() *CoursesDao {
	return &CoursesDao{
		group:   "default",
		table:   "courses",
		columns: coursesColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CoursesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CoursesDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CoursesDao) Columns() CoursesColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CoursesDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CoursesDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CoursesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
