package router

import (
	log "github.com/sirupsen/logrus"
	"github.com/dohr-michael/relationship/apis/app/services"
	"github.com/gin-gonic/gin"
)

var logCmd = log.WithFields(log.Fields{
	"module": "router",
})

func InitRouter(router *gin.Engine) {
	logCmd.Info("Init all routers")
	services.UniverseRouter("/universes", router)
	services.EntityRouter("/entities", router)
}
