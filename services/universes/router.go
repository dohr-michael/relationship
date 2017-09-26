package universes

import (
	"github.com/pressly/chi"
	"github.com/dohr-michael/relationship/services/tools/crud"
	log "github.com/sirupsen/logrus"
)

var logCmd = log.WithFields(log.Fields{
	"module": "universes.router",
})

var c = crud.Crud{
	Collection:          "universes",
	ItemsFactory:        func() crud.Entities { return &Universes{} },
	ItemFactory:         func() crud.Entity { return &Universe{} },
	ItemCreationFactory: func() crud.Entity { return &UniverseCreation{} },
	ItemUpdateFactory:   func() crud.Updatable { return &UniverseUpdate{} },
}

func Router(router chi.Router) {
	c.Router(router)
}
