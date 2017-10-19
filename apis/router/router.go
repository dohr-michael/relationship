package router

import (
	log "github.com/sirupsen/logrus"
	"github.com/dohr-michael/relationship/apis/app/services"
	"github.com/gin-gonic/gin"
	"github.com/dohr-michael/relationship/apis/tools"
)

var logCmd = log.WithFields(log.Fields{
	"module": "router",
})

func InitRouter(router *gin.Engine) {
	logCmd.Info("Init all routers")
	group := router.Group("/api")
	{
		group.Use(tools.NoCacheMiddleware)
		services.UniverseRouter("/universes", group)
		services.EntityRouter("/entities", group)
	}

}
