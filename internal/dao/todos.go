// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"NTNU_APP_2023_Final_Backend/internal/dao/internal"
)

// internalTodosDao is internal type for wrapping internal DAO implements.
type internalTodosDao = *internal.TodosDao

// todosDao is the data access object for table todos.
// You can define custom methods on it to extend its functionality as you wish.
type todosDao struct {
	internalTodosDao
}

var (
	// Todos is globally public accessible object for table todos operations.
	Todos = todosDao{
		internal.NewTodosDao(),
	}
)

// Fill with you ideas below.
