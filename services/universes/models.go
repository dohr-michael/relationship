package universes

import (
	"github.com/dohr-michael/relationship/services/tools/crud"
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
	"github.com/dohr-michael/relationship/services/tools/models"
)

type Universe struct {
	crud.BaseEntity `json:",inline" bson:",inline" valid:"required"`
	Name string     `json:"name" bson:"name" valid:"required"`
}

type Universes []Universe

func (u *Universes) Len() int {
	return len(*u)
}

type UniverseCreation struct {
	Universe `bson:",inline" valid:"required"`
}

func (u *UniverseCreation) UnmarshalJSON(data []byte) error {
	type alias UniverseCreation
	um := &struct{ *alias }{alias: (*alias)(u),}
	if err := json.Unmarshal(data, &um); err != nil {
		return err
	}
	um.Id = bson.NewObjectId()
	return nil
}

type UniverseUpdate struct {
	Name      string          `json:"name" bson:"name" valid:"required"`
	UpdatedAt models.DateTime `json:"-" bson:"updatedAt" valid:"-"`
}

func (u *UniverseUpdate) GetUpdatedAt() models.DateTime      { return u.UpdatedAt }
func (u *UniverseUpdate) SetUpdatedAt(value models.DateTime) { u.UpdatedAt = value }
