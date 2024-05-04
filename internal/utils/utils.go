package utils

import "math/rand"

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func CreateRandomString(length int) string {
	bytesArray := make([]byte, length)
	for i := range bytesArray {
		bytesArray[i] = letters[rand.Intn(len(letters))]
	}
	return string(bytesArray)
}
