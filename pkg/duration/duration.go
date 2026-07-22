package duration

import (
	"time"
)

func GetDuration(duration string) time.Duration {
	times, err := time.ParseDuration(duration)
	if err != nil || times <= 0 {
		return 5 * time.Minute
	}
	return times
}
