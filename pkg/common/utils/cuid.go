package utils

import (
	"crypto/rand"

	"github.com/lucsky/cuid"
)

func Cuid() string {

	c, err := cuid.NewCrypto(rand.Reader)
	if err != nil {

		return ""
	}

	return c
}
