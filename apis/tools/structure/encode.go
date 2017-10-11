package structure

import "encoding/json"

func Encode(obj interface{}) (map[string]interface{}, error) {
	if b, err := json.Marshal(obj); err != nil {
		return nil, err
	} else {
		res := make(map[string]interface{})
		err2 := json.Unmarshal(b, res)
		return res, err2
	}
}
