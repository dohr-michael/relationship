package crud

import (
	"net/http"
	"github.com/sirupsen/logrus"
	"github.com/dohr-michael/relationship/apis/tools"
	"github.com/dohr-michael/relationship/apis/tools/mongo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/gin-gonic/gin"
	"time"
)

var log = logrus.WithFields(logrus.Fields{
	"module": "tools.crud",
})

type Entities interface {
	Len() int
}

type WithEntity interface {
	SetEntity(entity *Entity)
	GetEntity() *Entity
}

type Entity struct {
	Id        bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty" binding:"-"`
	Index     int           `json:"index,omitempty" bson:"index,omitempty" binding:"-"`
	UpdatedBy string        `json:"updatedBy,omitempty" bson:"updatedBy,omitempty"`
	UpdatedAt time.Time     `json:"updatedAt,omitempty" bson:"updatedAt,omitempty" time_format:"2017-04-25T15:08:43.687Z"`
}

func (e *Entity) SetEntity(entity *Entity) {
	e.Id = entity.Id
	e.Index = entity.Index
	e.UpdatedBy = entity.UpdatedBy
	e.UpdatedAt = entity.UpdatedAt
}

func (e *Entity) GetEntity() *Entity {
	return e
}

type BaseEntities []Entity

type Crud struct {
	Collection          string
	ItemsFactory        func() Entities
	ItemFactory         func() WithEntity
	ItemCreationFactory func() WithEntity
	ItemUpdateFactory   func() WithEntity
}

func (c *Crud) Router(base string, router *gin.Engine) {
	log.Infof("Register %s", c.Collection)
	r := router.Group(base)
	{
		r.GET("/", c.Filter)
		r.GET("/:id", c.ById)
		r.POST("/", c.Create)
		r.PUT("/:id", c.Update)
		r.DELETE("/:id", c.Delete)
	}
}

type filterQuery struct {
	From int `form:"from"`
	Size int `form:"size"`
}

func (c *Crud) Filter(context *gin.Context) {
	// TODO Read props
	var query filterQuery
	context.Bind(&query)

	mongo.Col(c.Collection, func(col *mgo.Collection) {
		items := c.ItemsFactory()
		var total int
		var err error
		if total, err = col.Find(bson.M{}).Count(); err != nil {
			mongo.ToHttpError(err, context)
			return
		}
		var q = col.Find(bson.M{"index": bson.M{"$gte": query.From}}).Sort("-index")
		if query.Size != 0 {
			q = q.Limit(query.Size)
		}
		if err = q.All(items); err != nil {
			mongo.ToHttpError(err, context)
		}
		length := items.Len()
		res := &tools.Paginate{
			Length: length,
			Offset: query.From,
			Total:  total,
			Items:  items,
		}
		context.JSON(http.StatusOK, &res)
	})
}

func (c *Crud) ById(context *gin.Context) {
	id := context.Param("id")

	mongo.Col(c.Collection, func(col *mgo.Collection) {
		item := c.ItemFactory()
		if err := col.Find(bson.M{"_id": mongo.GetId(id),}).One(item); err != nil {
			mongo.ToHttpError(err, context, "byId", c.Collection, id)
			return
		}
		context.JSON(http.StatusOK, &item)
	})
}

func (c *Crud) Create(context *gin.Context) {
	body := c.ItemCreationFactory()
	log.Debug("Start creating...", c.Collection)
	if err := context.BindJSON(body); err != nil {
		log.Error("Error when reading payload", c.Collection, err.Error())
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	mongo.DB(func(db *mgo.Database) {
		nextIndex, err := mongo.GetNextIndex(c.Collection, db)
		if err != nil {
			mongo.ToHttpError(err, context, "create", c.Collection)
			return
		}
		col := db.C(c.Collection)
		// TODO UpdatedBy from context.
		body.SetEntity(&Entity{Id: bson.NewObjectId(), Index: nextIndex, UpdatedAt: time.Now(), UpdatedBy: "dohr.michael@gmail.com"})
		if err := col.Insert(&body); err != nil {
			mongo.ToHttpError(err, context, "create", c.Collection)
			return
		}
		context.JSON(http.StatusCreated, gin.H{"id": body.GetEntity().Id})
	})
}

func (c *Crud) Update(context *gin.Context) {
	id := context.Param("id")
	body := c.ItemUpdateFactory()
	if err := context.BindJSON(body); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	mongo.Col(c.Collection, func(col *mgo.Collection) {
		// TODO UpdatedBy from context.
		body.SetEntity(&Entity{UpdatedBy: "dohr.michael@gmail.com", UpdatedAt: time.Now()})
		if err := col.Update(bson.M{"_id": mongo.GetId(id)}, bson.M{"$set": body}); err != nil {
			mongo.ToHttpError(err, context, "update", c.Collection, id)
			return
		}
		res := c.ItemFactory()
		if err := col.Find(bson.M{"_id": mongo.GetId(id),}).One(res); err != nil {
			mongo.ToHttpError(err, context, "update", c.Collection, id)
			return
		}
		context.JSON(http.StatusCreated, &res)
	})
}

func (c *Crud) Delete(context *gin.Context) {

}
