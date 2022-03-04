package cron

import (
	"fmt"
	"time"
)

func demo() {
	time.Sleep(time.Second * 3)
	fmt.Println("hello world" + time.Now().Format("2006-01-02 15:04:05"))
}
