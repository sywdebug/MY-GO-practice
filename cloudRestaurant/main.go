package main

import (
	"cloudRestaurant/tool"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := tool.parseConfig("./config/app.json")
	if err != nil {
		panic(err.Error())
	}
	app := gin.Default()

	app.Run(cfg.AppHost + ":" + cfg.AppPort)
}
