package main

import (
	_ "hjmcloud/boot"
	_ "hjmcloud/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
