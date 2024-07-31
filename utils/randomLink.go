package utils

import "math/rand"

func RandomLink() string {
	abc := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	return randomString(abc, 6)
}

func randomString(abc string, i int) string {
	b := make([]byte, i)
	for i := range b {
		b[i] = abc[rand.Intn(len(abc))]
	}
	return string(b)
}
