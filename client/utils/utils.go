package utils

import (
	"os"
	"strconv"
)

func GetIntEnv(key string, fallback int) int {
	val := os.Getenv(key)
	ret, err := strconv.Atoi(val)
	if err != nil {
		return fallback
	}
	return ret
}
