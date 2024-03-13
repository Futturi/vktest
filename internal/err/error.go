package err

type Err struct {
	Error string `json:"error"`
}

func NewErr(err error) Err {
	return Err{Error: err.Error()}
}
