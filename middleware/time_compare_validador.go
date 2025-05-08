package middleware

import (
	"time"
)

func TimeCompareFutureValidator(startAt time.Time) bool {
	return startAt.After(time.Now())
}


func TimeComparePastValidator(endAt time.Time) bool {
	return endAt.Before(time.Now())
}

