package bootstrap

import (
	"fmt"
	"github.com/flc1125/event/app/model"
	"github.com/flc1125/event/pkg/database"
)

type Tests struct {
	ID   int
	Name string
}

type Article struct {
	ID    int
	Title string
	Text  string
}

func init() {
	// 配置 DB 连接
	database.Configure("default", &database.Config{
		Driver:   "mysql",
		Host:     "127.0.0.1",
		Port:     "3306",
		Username: "root",
		Password: "123456",
		Charset:  "utf8mb4",
		Database: "event",
	})

	database.Configure("1111", &database.Config{
		Driver:   "mysql",
		Host:     "192.168.8.1",
		Port:     "3306",
		Username: "1111",
		Password: "1111",
		Charset:  "latin1",
		Database: "1111",
	})

	var (
		tests Tests
		//article Article
	)
	database.Connect("default").DB.First(&tests)
	//fmt.Println(tests)

	//database.Connect("8591").DB.Table("article").First(&article)
	//fmt.Println(article)
	var t model.Tests
	database.Connect("default").DB.First(&t)
	fmt.Println(t)
}
