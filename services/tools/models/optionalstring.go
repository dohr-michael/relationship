package models

import (
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
	"fmt"
)

type OptionalString struct {
	Value *string
}

func (t OptionalString) MarshalJSON() ([]byte, error) {
	if t.Value == nil {
		return nil, nil
	}
	return json.Marshal(&t.Value)
}

func (t *OptionalString) UnmarshalJSON(data []byte) error {
	if data == nil {
		return nil
	}
	val := ""
	if err := json.Unmarshal(data, &val); err != nil {
		return err
	}
	t.Value = &val
	return nil
}

func (t OptionalString) GetBSON() (interface{}, error) {
	if t.Value == nil {
		return nil, nil
	}
	return *t.Value, nil
}

func (t *OptionalString) SetBSON(raw bson.Raw) error {
	if raw.Data == nil {
		return nil
	}
	val := ""
	if err := raw.Unmarshal(&val); err != nil {
		return err
	}
	t.Value = &val
	return nil
}

func (t OptionalString) String() string {
	if t.Value == nil {
		return "null"
	}
	return fmt.Sprintf("\"%s\"", *t.Value)
}
