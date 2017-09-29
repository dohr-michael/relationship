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

type Entity interface {
	GetId() bson.ObjectId
	GetIndex() int
	SetIndex(value int)
}

type Entities interface {
	Len() int
}

type BaseEntity struct {
	Id        bson.ObjectId `json:"id" bson:"_id" binding:"-"`
	Index     int           `json:"index" bson:"index" binding:"-"`
	UpdatedBy string        `json:"updatedBy" bson:"updatedBy"`
	UpdatedAt time.Time     `json:"updatedAt" bson:"updatedAt" time_format:"2017-04-25T15:08:43.687Z"`
}

func (e *BaseEntity) GetId() bson.ObjectId { return e.Id }
func (e *BaseEntity) GetIndex() int        { return e.Index }
func (e *BaseEntity) SetIndex(value int)   { e.Index = value }

type BaseEntities []BaseEntity

type Crud struct {
	Collection          string
	ItemsFactory        func() Entities
	ItemFactory         func() Entity
	ItemCreationFactory func() Entity
	ItemUpdateFactory   func() interface{}
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
	if err := context.BindJSON(body); err != nil {
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
		body.SetIndex(nextIndex)
		if err := col.Insert(body); err != nil {
			mongo.ToHttpError(err, context, "create", c.Collection)
			return
		}
		context.JSON(http.StatusCreated, gin.H{"id": body.GetId()})
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
