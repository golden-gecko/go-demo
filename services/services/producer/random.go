package main

import "math/rand"

func RandomPlate() string {
	const characters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const digits = "0123456789"

	return RandomString(4, characters) + RandomString(8, digits)
}

func RandomString(length int, charset string) string {
	b := make([]byte, length)

	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}

	return string(b)
}

func RandomWord(length int) string {
	const characters = "abcdefghijklmnopqrstuvwxyz"

	return RandomString(length, characters)
}

func RandomYear() int {
	return 1980 + rand.Intn(30)
}
