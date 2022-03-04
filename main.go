package main

import (
	_ "github.com/flc1125/event/bootstrap"
	"github.com/flc1125/event/cron"
)

func main() {
	//cmd.Execute()
	cron.Run()
}
