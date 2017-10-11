package models

type Universe struct {
	Hash string `json:"hash" binding:"-"`
	Name string `json:"name" binding:"required"`
}

type Universes []Universe

func (u *Universes) Len() int {
	return len(*u)
}

type UniverseUpdate struct {
	Name string `json:"name" bson:"name" binding:"required"`
}
