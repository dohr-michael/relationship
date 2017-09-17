package universes

import (
	"github.com/pressly/chi"
	"net/http"
	"github.com/dohr-michael/relationship/services/tools"
	log "github.com/sirupsen/logrus"
)

var logCmd = log.WithFields(log.Fields{
	"module": "universes.router",
})

func Router(router chi.Router) {
	logCmd.Debug("Initialize Router")
	router.Post("/", filter)
	router.Get("/:id", byId)
}

func filter(w http.ResponseWriter, r *http.Request) {
	body := tools.SearchRequest{}
	tools.DecodeJson(&body, r)

	tools.JsonResult(&tools.Paginate{
		Length: 2,
		Offset: 0,
		Total:  2,
		Items: &Universes{
			Universe{"1", "Vampire : L.A. By Night"},
			Universe{"2", "Star Wars : La Règle des deux"},
		},
	})(w, r)
}

func byId(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	tools.JsonResult(&Universe{id, "Vampire : L.A. By Night"})(w, r)
}
