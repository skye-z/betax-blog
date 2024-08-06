package util

import (
	"fmt"
	"log"
)

func OutLog(module, msg string) {
	log.Println("["+module+"]", msg)
}

func OutLogf(module, format string, v ...any) {
	log.Println("["+module+"]", fmt.Sprintf(format, v...))
}

func OutErr(module, format string, v ...any) {
	log.Fatalf("["+module+"]", fmt.Sprintf(format, v...))
}
