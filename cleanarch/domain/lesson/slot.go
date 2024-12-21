package lesson

import (
	"time"
)

type Slot struct {
	tutorID string
	startAt time.Time
	endAt time.Time
}

func NewSlot(tutorID string, startAt time.Time, endAt time.Time) *Slot {
	return &Slot {
		tutorID: tutorID,
		startAt: startAt,
		endAt: endAt,
	}
}