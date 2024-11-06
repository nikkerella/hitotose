package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	handler "github.com/nikkerella/hitotose/gin/api"
	mongo "github.com/nikkerella/hitotose/gin/db/mongo"
	"github.com/nikkerella/hitotose/gin/db/redis"
)

func main() {
	// Databases
	redis.Cli = redis.NewRedisClient()
	mongo.Init()

	// Router
	router := gin.Default()
	router.Use(cors.Default())
	handler.Init(router)
	router.Run(":8080")
}
