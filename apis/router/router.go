package router

import (
	"github.com/pressly/chi"
	log "github.com/sirupsen/logrus"
	"github.com/dohr-michael/relationship/apis/app/services"
)

var logCmd = log.WithFields(log.Fields{
	"module": "router",
})

func InitRouter(router *chi.Mux) {
	services.UniverseRouter("/universes", router)
}
