package shopownerdomain

import (
	"time"
)

type Reservation struct {
	id int64
	reservationDatetime time.Time
	status string
}