package cmd

import (
	"NTNU_APP_2023_Final_Backend/utility"
	"github.com/gogf/gf/v2/frame/g"
)

func initTablesData() {
	initUsers()
}

func initUsers() {
	m := g.DB().Model("users")
	encodedHash, _ := utility.HashPWD("1234567890")
	_, err := m.Insert(g.Map{
		"email":    "admin@ntnu.app",
		"password": encodedHash,
	})
	utility.IfErrExit(err)
}
