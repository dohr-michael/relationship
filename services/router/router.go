package router

import (
	"github.com/pressly/chi"
	log "github.com/sirupsen/logrus"
	"github.com/dohr-michael/relationship/services/universes"
)

var logCmd = log.WithFields(log.Fields{
	"module": "router",
})

func InitRouter(router *chi.Mux) {
	router.Route("/universes", universes.Router)
}
