package utils

import "math/rand"

func RandomInt(n int) int {
	return rand.Intn(n)
}

func RandomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[RandomInt(len(letters))]
	}
	return string(b)
}
