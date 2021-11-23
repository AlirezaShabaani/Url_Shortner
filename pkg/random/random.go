package random

import (
	"math/rand"
)

const letterBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

type ShortenerService interface {
	New(lenght int) string
}

type shortenerService struct{}

func New() ShortenerService {
	return &shortenerService{}
}
func (s shortenerService) New(length int) string {
	return RandomString(length)
}

func RandomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}