package structure

import (
	"testing"
	"strings"
	"github.com/magiconair/properties/assert"
)

type Inner struct {
	Id   string `json:"id"`
	Name string `json:"name,omitempty"`
}

type TestStruct struct {
	Inner                         `json:",inline"`
	Field1 string                 `json:"field1"`
	Field2 bool                   `json:"field2"`
	Field3 int64                  `json:"field3"`
	Field4 float64                `json:"field4"`
	Field5 []string               `json:"field5"`
	Field6 map[string]interface{} `json:"field6"`
}

func TestDecodeError(t *testing.T) {
	obj := TestStruct{}
	mapObj := map[string]interface{}{
		"id":     "",
		"name":   "",
		"field1": "",
		"field2": "",
		"field3": "",
		"field4": "",
		"field5": "",
		"field6": "",
	}
	err := Decode(&obj, mapObj)
	if err == nil {
		t.Fatal("Unmarshal must be an error")
	}
	if !strings.Contains(err.Error(), "struct field TestStruct.field2 of type bool") {
		t.Fatal("Bad error message", err.Error())
	}
}

// Data must be initialized with default values.
func TestDecodeWithOmit(t *testing.T) {
	obj := TestStruct{}
	mapObj := map[string]interface{}{}
	err := Decode(&obj, mapObj)
	if err != nil {
		t.Fatal("Unmarshal must populate the object without error", err)
	}
}

func TestDecode(t *testing.T) {
	obj := TestStruct{}
	mapObj := map[string]interface{}{
		"id":     "42",
		"name":   "dent",
		"field1": "Hello",
		"field2": true,
		"field3": 42,
		"field4": 66.6,
		"field5": []string{"a", "b"},
		"field6": map[string]interface{}{"a": "a"},
	}
	err := Decode(&obj, mapObj)
	if err != nil {
		t.Fatal("Unmarshal must populate the object without error", err)
	}
	assert.Equal(t, obj.Id, "42")
	assert.Equal(t, obj.Name, "dent")
	assert.Equal(t, obj.Field1, "Hello")
	assert.Equal(t, obj.Field2, true)
	assert.Equal(t, obj.Field3, int64(42))
	assert.Equal(t, obj.Field4, float64(66.6))
	assert.Equal(t, len(obj.Field5), 2)
	assert.Equal(t, obj.Field5[0], "a")
	assert.Equal(t, obj.Field5[1], "b")
	assert.Equal(t, obj.Field6["a"], "a")
}
