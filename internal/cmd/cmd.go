package cmd

import (
	"NTNU_APP_2023_Final_Backend/utility"
	"context"
	"database/sql"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind()
			})
			s.Run()
			return nil
		},
	}
)

func InitDB() {
	// init db
	db, err := sql.Open("sqlite3", "./NTNU_APP_2023_Final_DB.sqlite3")
	utility.IfErrExit(err)
	defer db.Close()

	// init tables
	sqlFilename := []string{
		"users.sql",
	}
	for _, filename := range sqlFilename {
		sqlFile, err := os.ReadFile("./manifest/sql/" + filename)
		utility.IfErrExit(err)
		_, err = db.Exec(string(sqlFile))
		utility.IfErrExit(err)
	}
	fmt.Println("init db & tables success")

	// insert default data
	initTablesData()
	fmt.Println("insert default data success")
}
