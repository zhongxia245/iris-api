// Package main shows how an orm can be used within your web app
// it just inserts a column and select the first.

// BUG: rizla 热更新，在代码报错后，就无法重启
// rizla 作者没有修复这个问题，暂时弃用，改用 gowatch

package main

import (
	"time"

	"github.com/kataras/iris"

	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
)

/*
	go get -u github.com/mattn/go-sqlite3
	go get -u github.com/go-xorm/xorm

	If you're on win64 and you can't install go-sqlite3:
		1. Download: https://sourceforge.net/projects/mingw-w64/files/latest/download
		2. Select "x86_x64" and "posix"
		3. Add C:\Program Files\mingw-w64\x86_64-7.1.0-posix-seh-rt_v5-rev1\mingw64\bin
		to your PATH env variable.

	Docs: http://xorm.io/docs/
*/

// User is our user table structure.
type User struct {
	ID        int64  // auto-increment by-default by xorm
	Version   string `xorm:"varchar(200)"`
	Salt      string
	Username  string
	Password  string    `xorm:"varchar(200)"`
	Languages string    `xorm:"varchar(200)"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
}

func main() {
	app := iris.New()

	// xorm 的文档使用 http://gobook.io/read/github.com/go-xorm/manual-zh-CN/chapter-07/1.deleted.html
	orm, err := xorm.NewEngine("sqlite3", "./test.db")
	if err != nil {
		app.Logger().Fatalf("orm failed to initialized: %v", err)
	}

	iris.RegisterOnInterrupt(func() {
		orm.Close()
	})

	err = orm.Sync2(new(User))

	if err != nil {
		app.Logger().Fatalf("orm failed to initialized User table: %v", err)
	}

	// 插入数据
	app.Get("/insert", func(ctx iris.Context) {
		user := &User{Username: "kataras", Salt: "hash---", Password: "hashed", CreatedAt: time.Now(), UpdatedAt: time.Now()}
		orm.Insert(user)
		ctx.JSON(iris.Map{"data": user, "status": 0})
	})

	app.Post("/insert", func(ctx iris.Context) {
		var user User
		ctx.ReadJSON(&user)
		orm.Insert(user)
		ctx.JSON(iris.Map{"data": user, "status": 0})
	})

	// 获取单条数据
	app.Get("/get/{id:int64}", func(ctx iris.Context) {
		id, err := ctx.Params().GetInt64("id")
		if err != nil {
			app.Logger().Fatalf("orm failed to get url params: %v", err)
		}
		user := User{ID: id}
		if ok, _ := orm.Get(&user); ok {
			ctx.JSON(iris.Map{"data": user, "status": 0})
		}
	})

	// 删除数据
	app.Get("/delete/{id:int64}", func(ctx iris.Context) {
		id, err := ctx.Params().GetInt64("id")
		if err != nil {
			app.Logger().Fatalf("orm failed to get url params: %v", err)
		}
		user := User{ID: id}
		orm.Delete(&user)
		ctx.JSON(iris.Map{"data": true, "message": "删除成功", "status": 0})
	})

	// 获取所有数据的列表
	app.Get("/list", func(ctx iris.Context) {
		users := make([]User, 0)
		err := orm.Limit(20, 0).Find(&users)

		if err != nil {
			app.Logger().Fatalf("orm failed to get list data: %v", err)
		}
		ctx.JSON(iris.Map{"data": users, "status": 0})
	})

	// http://localhost:10086/insert
	// http://localhost:10086/get
	app.Run(iris.Addr(":10086"), iris.WithoutServerError(iris.ErrServerClosed))
}
