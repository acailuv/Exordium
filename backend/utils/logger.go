package utils

import (
	"fmt"
	"log"
	"runtime"
)

func GetCurrentCodePosition() string {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		return fmt.Sprintf("(%s) <Line %d>", file, line)
	}
	return "Unknown"
}

func LogError(where string, err error, data ...interface{}) {
	log.SetPrefix("[ERROR]\t")
	log.Printf("%v: %v: %v", where, err, data)
}

func LogWarning(where string, data ...interface{}) {
	log.SetPrefix("[WARNING]\t")
	log.Printf("%v: %v", where, data)
}

func LogInfo(where string, data ...interface{}) {
	log.SetPrefix("[INFO]\t")
	log.Printf("%v: %v", where, data)
}
