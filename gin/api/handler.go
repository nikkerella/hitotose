package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nikkerella/hitotose/gin/api/game"
	"github.com/nikkerella/hitotose/gin/api/rank"
	"github.com/nikkerella/hitotose/gin/api/user"
)

func Init(e *gin.Engine) {
	assets(e)

	game.Init(e)
	rank.Init(e)
	user.Init(e)
}

func assets(e *gin.Engine) {
	e.Static("/assets", "./assets")
}
