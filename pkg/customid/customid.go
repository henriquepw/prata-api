package customid

import gonanoid "github.com/matoous/go-nanoid/v2"

func createID(size int) string {
	return gonanoid.MustGenerate("123456789ABCDEFGHIJKLMNPQRSTUVWXYZ", size)
}

func New() string {
	return createID(18)
}

func NewTiny() string {
	return createID(6)
}
