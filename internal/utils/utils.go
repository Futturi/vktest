package utils

import (
	"crypto/sha1"
	"fmt"
)

const (
	salt = "erighwdjslaksjweqwojke2pqijfklbdfjnkjasnlkqergknlwenfkjeoiq2jflkgjevwnlkjqwn"
)

func HashPass(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
