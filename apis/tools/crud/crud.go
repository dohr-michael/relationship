package crud

import (
	"net/http"
	"github.com/sirupsen/logrus"
	"github.com/dohr-michael/relationship/apis/tools"
	"github.com/gin-gonic/gin"
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
	Hash  string `json:"hash,omitempty" binding:"-"`
	Index int    `json:"index,omitempty" binding:"-"`
}

func (e *Entity) SetEntity(entity *Entity) {
	e.Hash = entity.Hash
	e.Index = entity.Index
}

func (e *Entity) GetEntity() *Entity {
	return e
}

type BaseEntities []Entity

type Cru interface {
	Filter(context *gin.Context)
	ById(context *gin.Context)
	Create(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type Crud struct {
	Type                string
	ItemsFactory        func() Entities
	ItemFactory         func() WithEntity
	ItemCreationFactory func() WithEntity
	ItemUpdateFactory   func() WithEntity
}

func (c *Crud) Router(base string, router *gin.Engine) {
	log.Infof("Register %s", c.Type)
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
	res := &tools.Paginate{
		Size: 0,
		Total:  0,
		Page: 0,
		Items:  make([]interface{}, 0),
	}
	var err error
	if err != nil {

		return
	}
	context.JSON(http.StatusOK, &res)
}

func (c *Crud) ById(context *gin.Context) {
	//hash := context.Param("hash")
	item := c.ItemFactory()
	var err error
	if err != nil {
		return
	}
	context.JSON(http.StatusOK, &item)
}

func (c *Crud) Create(context *gin.Context) {
	body := c.ItemCreationFactory()
	log.Debug("Start creating...", c.Type)
	if err := context.BindJSON(body); err != nil {
		log.Error("Error when reading payload", c.Type, err.Error())
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var err error
	if err != nil {
		return
	}
	context.JSON(http.StatusCreated, gin.H{"hash": body.GetEntity().Hash})
}

func (c *Crud) Update(context *gin.Context) {
	//hash := context.Param("hash")
	body := c.ItemUpdateFactory()
	if err := context.BindJSON(body); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res := c.ItemFactory()
	var err error
	if err != nil {
		return
	}
	context.JSON(http.StatusCreated, &res)
}

func (c *Crud) Delete(context *gin.Context) {
	//hash := context.Param("hash")
}
