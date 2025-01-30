package common

import "syscall"

func GetEnv(key, fallback string) string {
	value, ok := syscall.Getenv(key)
	if !ok {
		return fallback
	}
	return value
}
