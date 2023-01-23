package main

import (
	"os"
)

func GetEnvironmentVariable(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
