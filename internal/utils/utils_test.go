package utils

import (
	"crypto/sha1"
	"fmt"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestUtils_Utils(t *testing.T) {
	hash := sha1.New()
	hash.Write([]byte("password"))
	password := HashPass("password")
	assert.Equal(t, fmt.Sprintf("%x", hash.Sum([]byte(salt))), password)
}
