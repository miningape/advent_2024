package util

import (
	"fmt"
	"time"
)



func MeasureRuntime(start time.Time, messages ...string) {
	elapsed := time.Since(start)
	message := ""
	for _, m := range messages {
		message += " " + m
	}

	fmt.Println(message, "- Took:", elapsed)
}