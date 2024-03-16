package errs

import "encoding/json"

type Err struct {
	Error string `json:"error"`
}

func NewErr(err error) []byte {
	rr := Err{Error: err.Error()}
	byt, _ := json.Marshal(rr)
	return byt
}
