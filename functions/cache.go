package functions

import (
	"fmt"
	"time"
)

func addTime(duration_string string) string {
	duration, _ := time.ParseDuration(duration_string)
	t := time.Now().Add(duration)
	fmt.Println(t.Format(time.RFC3339))
	return ""
}
