package router

import (
	"github.com/pressly/chi"
	log "github.com/sirupsen/logrus"
)

var logCmd = log.WithFields(log.Fields{
	"module": "router",
})

func InitRouter(router *chi.Mux) {
}
