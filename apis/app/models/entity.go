package models

import "github.com/dohr-michael/relationship/apis/tools/crud"

type Entity struct {
	crud.Entity
	Category string `json:"category" bson:"category" binding:"required"`
}
