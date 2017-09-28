package tools

import (
	"encoding/json"
	"github.com/asaskevich/govalidator"
	"log"
	"net/http"
)

func DecodeJson(res interface{}, req *http.Request) {
	dec := json.NewDecoder(req.Body)
	defer req.Body.Close()
	if err := dec.Decode(&res); err != nil {
		Panic(err.Error(), 500)
	}
	ok, err := govalidator.ValidateStruct(res)
	if !ok {
		Panic(err.Error(), 400)
	}
}

func JsonResult(res interface{}) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if res != nil {
			bytes, err := json.Marshal(&res)
			if err != nil {
				Panic(err.Error(), 500)
			} else {
				w.Header().Set("Content-Type", "application/json")
				w.Write(bytes)
			}
		} else {
			log.Println("No Content ...")
		}
	}
}
