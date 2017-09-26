package crud

import (
	"github.com/pressly/chi"
	"net/http"
	log "github.com/sirupsen/logrus"
	"github.com/dohr-michael/relationship/services/tools"
	"github.com/dohr-michael/relationship/services/tools/mongo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"github.com/dohr-michael/relationship/services/tools/models"
)

var logCmd = log.WithFields(log.Fields{
	"module": "tools.crud",
})

type Updatable interface {
	GetUpdatedAt() models.DateTime
	SetUpdatedAt(value models.DateTime)
}

type Entity interface {
	Updatable
	GetId() bson.ObjectId
	GetIndex() int
	SetIndex(value int)
}

type Entities interface {
	Len() int
}

type BaseEntity struct {
	Id        bson.ObjectId   `json:"id" bson:"_id" valid:"-"`
	Index     int             `json:"index" bson:"index" valid:"-"`
	UpdatedAt models.DateTime `json:"updatedAt" bson:"updatedAt" valid:"-"`
}

func (e *BaseEntity) GetId() bson.ObjectId               { return e.Id }
func (e *BaseEntity) GetIndex() int                      { return e.Index }
func (e *BaseEntity) SetIndex(value int)                 { e.Index = value }
func (e *BaseEntity) GetUpdatedAt() models.DateTime      { return e.UpdatedAt }
func (e *BaseEntity) SetUpdatedAt(value models.DateTime) { e.UpdatedAt = value }

type BaseEntities []BaseEntity

type Crud struct {
	Collection          string
	ItemsFactory        func() Entities
	ItemFactory         func() Entity
	ItemCreationFactory func() Entity
	ItemUpdateFactory   func() Updatable
}

func (c *Crud) Router(router chi.Router) {
	logCmd.Infof("Register %s", c.Collection)
	router.Get("/", c.Filter)
	router.Get("/{id}", c.ById)
	router.Post("/", c.Create)
	router.Put("/{id}", c.Update)
	router.Delete("/{id}", c.Delete)
}

func (c *Crud) Filter(w http.ResponseWriter, r *http.Request) {
	// TODO Read props
	var from int = 0
	if pFrom, err := strconv.ParseInt(r.URL.Query().Get("from"), 10, 64); err == nil {
		from = int(pFrom)
	}
	var size int = 0
	if pSize, err := strconv.ParseInt(r.URL.Query().Get("size"), 10, 64); err == nil {
		size = int(pSize)
	}
	mongo.Col(c.Collection, func(col *mgo.Collection) {
		items := c.ItemsFactory()
		total, err1 := col.Find(bson.M{}).Count()
		mongo.ParseError(err1)
		var q = col.Find(bson.M{"index": bson.M{"$gte": from}}).Sort("-index")
		if size != 0 {
			q = q.Limit(size)
		}
		mongo.ParseError(q.All(items))
		length := items.Len()
		tools.JsonResult(&tools.Paginate{
			Length: length,
			Offset: from,
			Total:  total,
			Items:  items,
		})(w, r)
	})
}

func (c *Crud) ById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	mongo.Col(c.Collection, func(col *mgo.Collection) {
		item := c.ItemFactory()
		mongo.ParseError(col.Find(bson.M{"_id": mongo.GetId(id),}).One(item), c.Collection, id)
		tools.JsonResult(&item)(w, r)
	})
}

func (c *Crud) Create(w http.ResponseWriter, r *http.Request) {
	body := c.ItemCreationFactory()
	tools.DecodeJson(body, r)
	mongo.DB(func(db *mgo.Database) {
		nextIndex := mongo.GetNextIndex(c.Collection, db)
		col := db.C(c.Collection)
		body.SetIndex(nextIndex)
		body.SetUpdatedAt(models.NewDateTime())
		mongo.ParseError(col.Insert(body))
		tools.JsonResult(map[string]interface{}{
			"id": body.GetId(),
		})(w, r)
	})
}

func (c *Crud) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	body := c.ItemUpdateFactory()
	tools.DecodeJson(body, r)
	body.SetUpdatedAt(models.NewDateTime())
	mongo.Col(c.Collection, func(col *mgo.Collection) {
		mongo.ParseError(col.Update(bson.M{"_id": mongo.GetId(id)}, bson.M{"$set": body}))
		res := c.ItemFactory()
		mongo.ParseError(col.Find(bson.M{"_id": mongo.GetId(id),}).One(res))
		tools.JsonResult(&res)(w, r)
	})
}

func (c *Crud) Delete(w http.ResponseWriter, r *http.Request) {

}
