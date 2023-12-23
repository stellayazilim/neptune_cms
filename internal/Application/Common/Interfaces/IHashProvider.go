package Interfaces

type IHashProvider interface {
	GenHash(str []byte) []byte
	Compare([]byte, []byte) bool
}
