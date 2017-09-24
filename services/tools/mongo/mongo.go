package mongo

import (
	"gopkg.in/mgo.v2"
	tools "github.com/dohr-michael/relationship/services/tools"
	"github.com/dohr-michael/relationship/services/cfg"
)

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