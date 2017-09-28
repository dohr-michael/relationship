package tools

import "encoding/json"

type Error struct {
	Message string
	Code    int
}

func (err *Error) Error() string {
	return err.Message
}

type mess struct {
	Error   string        `json:"error"`
	Details []interface{} `json:"details"`
}

func Panic(message string, code int, args ...interface{}) {
	var errorMessage string = ""
	m := mess{Error: message, Details: args}
	b, err := json.Marshal(m)
	if err == nil {
		errorMessage = string(b)
	}
	panic(Error{
		Code:    code,
		Message: errorMessage,
	})
}
