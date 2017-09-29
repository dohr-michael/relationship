package mongo

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/dohr-michael/relationship/apis/cfg"
	"github.com/gin-gonic/gin"
	"net/http"
)

var indexCol = "colIdx"

type collectionIndex struct {
	Id         string `bson:"_id"`
	Collection string `bson:"collection"`
	Index      int    `bson:"index"`
}

func GetNextIndex(col string, db *mgo.Database) (int, error) {
	res := collectionIndex{}
	collection := db.C(indexCol)
	if _, err := collection.Upsert(bson.M{"collection": col,}, bson.M{"$set": bson.M{"collection": col},}); err != nil {
		return 0, err
	}
	change := mgo.Change{Update: bson.M{"$inc": bson.M{"index": 1,},}, ReturnNew: false}
	if _, err := collection.Find(bson.M{"collection": col,}).Sort("-index").Apply(change, &res); err != nil {
		return 0, err
	}
	return res.Index, nil
}

func GetId(id string) interface{} {
	if bson.IsObjectIdHex(id) {
		return bson.ObjectIdHex(id)
	}
	return id
}

func ToHttpError(err error, c *gin.Context, details ...interface{}) {
	errorPayload := gin.H{"error": err.Error(), "details": details}
	if err == mgo.ErrNotFound {
		c.JSON(http.StatusNotFound, errorPayload)
	} else {
		c.JSON(http.StatusInternalServerError, errorPayload)
	}
}

func DB(fn func(db *mgo.Database)) {
	session, err := mgo.Dial(cfg.GetMongoUrl())
	if err != nil {
		panic(err)
	}
	defer session.Close()
	db := session.DB(cfg.GetMongoDatabase())
	fn(db)
}

func Col(col string, fn func(collection *mgo.Collection)) {
	DB(func(db *mgo.Database) {
		c := db.C(col)
		fn(c)
	})
}
