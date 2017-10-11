package structure

import (
	"encoding/json"
)

func Decode(obj interface{}, m map[string]interface{}) error {
	if b, err := json.Marshal(m); err != nil {
		return err
	} else {
		return json.Unmarshal(b, obj)
	}
	/*
	value := reflect.ValueOf(obj).Elem()
	for k, v := range m {
		f := value.FieldByName(k)
		if !f.IsValid() {
			return fmt.Errorf("No such field: %s in obj.", k)
		}
		if !f.CanSet() {
			return fmt.Errorf("Cannot set %s field value.", k)
		}
		fType := f.Type()
		val := reflect.ValueOf(v)
		if fType != val.Type() {
			return fmt.Errorf("The type of %s dont match with the expected type of the struc.", k)
		}
		f.Set(val)
	}
	return nil
	*/
}

