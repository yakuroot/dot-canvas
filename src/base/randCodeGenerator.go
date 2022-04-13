package base

import "math/rand"

func GetRandCode() string {
	letter := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	code := make([]rune, 18)

	for i := range code {
		code[i] = letter[rand.Intn(len(letter))]
	}

	return string(code)
}
