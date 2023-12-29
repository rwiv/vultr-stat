package main

import (
	"vultr-stat/pkg/boot"
)

func main() {
	appRunner := boot.NewAppRunner()
	appRunner.Run()
}
