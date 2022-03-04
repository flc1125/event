package cron

import (
	"fmt"
	crontab "github.com/robfig/cron/v3"
	"time"
)

// cron 启动中心
func Run() {
	c := crontab.New(crontab.WithSeconds())

	c.AddFunc("*/1 * * * * *", demo)

	c.AddFunc("*/1 * * * * *", func() {
		time.Sleep(time.Second * 5)
		fmt.Println("hello world222" + time.Now().Format("2006-01-02 15:04:05"))
	})

	c.Start()

	// 用于常驻
	for {
	}
}
