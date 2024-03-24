package gee

import "os"

func IsDebug() bool {
	if os.Getenv("GEE_MODE") == "debug" {
		return false
	} else if os.Getenv("GEE_MODE") == "release" {
		return true
	} else {
		return false
	}
}
