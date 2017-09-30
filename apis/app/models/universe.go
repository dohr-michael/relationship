package models

import (
	"github.com/dohr-michael/relationship/apis/tools/crud"
)

type Universe struct {
	crud.Entity `json:",inline" bson:",inline" binding:"-"`
	Name string `json:"name" bson:"name" binding:"-"`
}

type Universes []Universe

func (u *Universes) Len() int {
	return len(*u)
}

type UniverseCreation struct {
	crud.Entity `json:",inline" bson:",inline" binding:"-"`
	Name string `json:"name" bson:"name" binding:"required"`
}

type UniverseUpdate struct {
	crud.Entity `json:",inline" bson:",inline" binding:"-"`
	Name string `json:"name" bson:"name" binding:"required"`
}
