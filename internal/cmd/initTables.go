package cmd

import (
	"NTNU_APP_2023_Final_Backend/internal/consts"
	"NTNU_APP_2023_Final_Backend/utility"
	"github.com/gogf/gf/v2/frame/g"
)

func initTablesData() {
	initUsers()
	initCourses()
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

func initCourses() {
	insertList := g.List{}
	for _, course := range consts.Courses {
		insertList = append(insertList, g.Map{
			"number":   course.Number,
			"code":     course.Code,
			"name":     course.Name,
			"teacher":  course.Teacher,
			"time":     course.Time,
			"location": course.Location,
		})
	}
	_, err := g.Model("courses").Data(insertList).Insert()
	utility.IfErrExit(err)
}
