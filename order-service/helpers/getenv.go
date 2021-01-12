package helpers

import "os"

var cache = make(map[string]string)

// Getenv function increase performance via using memory cache...
func Getenv(key string) string {
	if v, ok := cache[key]; ok {
		return v
	}
	cache[key] = os.Getenv(key)
	return cache[key]
}
