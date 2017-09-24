package universes

import (
	"github.com/pressly/chi"
	"net/http"
	"github.com/dohr-michael/relationship/services/tools"
	"github.com/dohr-michael/relationship/services/tools/mongo"
	log "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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

	mongo.Col("universes", func(col *mgo.Collection) {
		items := Universes{}
		// TODO Pagination.
		mongo.ParseError(col.Find(bson.D{}).All(&items))
		tools.JsonResult(&tools.Paginate{
			Length: len(items),
			Offset: 0,
			Total: len(items),
			Items: &items,
		})(w, r)
	})
}

func byId(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	tools.JsonResult(&Universe{id, "Vampire : L.A. By Night"})(w, r)
}
