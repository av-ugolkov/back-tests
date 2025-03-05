//go:build debug

package logger

import "log"

func Log(msg string) {
	log.Println("[DEBUG]:", msg)
}
