package Providers

import (
	"github.com/stellayazilim/neptune.application/Common/Interfaces"
	"golang.org/x/crypto/bcrypt"
)

type hashProvider struct{}

func HashProvider() Interfaces.IHashProvider {
	return &hashProvider{}
}

func (h *hashProvider) GenHash(str []byte) []byte {

	bytes, err := bcrypt.GenerateFromPassword(str, 10)

	if err != nil {
	}
	return bytes
}

func (h *hashProvider) Compare(plain []byte, hash []byte) bool {

	if err := bcrypt.CompareHashAndPassword(hash, plain); err != nil {
		return false
	}
	return true
}
