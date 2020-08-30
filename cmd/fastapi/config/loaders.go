package config

import (
	"os"
	"strconv"
	"strings"
)

//Get environment config as string value
func getString(key, defaultVal string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultVal
	}
	return val
}

//Get environment config as integer value
func getInt(key string, defaultVal int) int {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultVal
	}
	intVal, err := strconv.Atoi(val)
	if err != nil {
		return defaultVal
	}
	return intVal
}

//Get environment config as boolean value
func getBool(key string, defaultVal bool) bool {
	val, ok := os.LookupEnv(key)
	if ok && val == "true" {
		return true
	}
	return false
}

//Get environment config as slice value
func getSlice(key string, defaultVal []string) []string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultVal
	}
	return strings.Split(val, ",")
}
