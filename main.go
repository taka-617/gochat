package main

import (
	"github.com/taka-617/gochat/db"
	"github.com/taka-617/gochat/route"
)

func main() {
	db.Init()

	web := route.GetRouter()
	web.Run(":80")
}
