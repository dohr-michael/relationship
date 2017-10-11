package models

import "github.com/dohr-michael/relationship/apis/tools/crud"

type Entity struct {
	crud.Entity                       `json:",inline" binding:"-"`
	Category   string                 `json:"category" binding:"required"`
	Name       string                 `json:"name" binding:"required"`
	Properties map[string]interface{} `json:"properties" binding:"required"`
}

type Entities []Entity

func (e *Entities) Len() int { return len(*e) }

type EntityUpdate struct {
	crud.Entity                       `json:",inline" binding:"-"`
	Name       string                 `json:"name,omitempty" binding:"-"`
	Properties map[string]interface{} `json:"properties,omitempty" binding:"-"`
}
