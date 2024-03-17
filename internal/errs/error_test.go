package errs

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestErrs_NewErr(t *testing.T) {
	rr := Err{Error: "error"}
	byt, err := json.Marshal(rr)
	assert.Equal(t, byt, NewErr(errors.New("error")))
	assert.Equal(t, err, nil)
}
