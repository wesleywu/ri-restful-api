package main

import (
	_ "github.com/WesleyWu/ri-restful-api/boot"
	_ "github.com/WesleyWu/ri-restful-api/router"
	"github.com/gogf/gf/v2/frame/g"
)

func main() {
	s := g.Server()
	s.Run()
}
