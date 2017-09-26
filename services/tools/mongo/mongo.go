package mongo

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/dohr-michael/relationship/services/tools"
	"github.com/dohr-michael/relationship/services/cfg"
)

var indexCol = "colIdx"

type collectionIndex struct {
	Id         string `bson:"_id"`
	Collection string `bson:"collection"`
	Index      int    `bson:"index"`
}

func GetNextIndex(col string, db *mgo.Database) int {
	res := collectionIndex{}
	collection := db.C(indexCol)
	_, err := collection.Upsert(bson.M{"collection": col,}, bson.M{"$set": bson.M{"collection": col},})
	ParseError(err, "GetNextIndex", col)
	change := mgo.Change{Update: bson.M{"$inc": bson.M{"index": 1,},}, ReturnNew: false}
	_, err2 := collection.Find(bson.M{"collection": col,}).Sort("-index").Apply(change, &res)
	ParseError(err2, "GetNextIndex", col)
	return res.Index
}

func GetId(id string) interface{} {
	if bson.IsObjectIdHex(id) {
		return bson.ObjectIdHex(id)
	}
	return id
}

func ParseError(err error, details ...interface{}) {
	if err != nil {
		if err == mgo.ErrNotFound {
			tools.Panic("not.found", 404, details...)
		} else {
			tools.Panic(err.Error(), 500, details...)
		}
	}
}

func DB(fn func(db *mgo.Database)) {
	session, err := mgo.Dial(cfg.GetMongoUrl())
	if err != nil {
		ParseError(err)
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
