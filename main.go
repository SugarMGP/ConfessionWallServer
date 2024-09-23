package main

import (
	"ConfessionWall/app/midwares"
	"ConfessionWall/config/database"
	"ConfessionWall/config/router"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Init()
	r := gin.Default()
	r.NoMethod(midwares.HandleNotFound) // 中间件统一处理404
	r.NoRoute(midwares.HandleNotFound)
	router.Init(r)

	err := r.Run()
	if err != nil {
		log.Fatal("Server start failed: ", err)
	}
}
