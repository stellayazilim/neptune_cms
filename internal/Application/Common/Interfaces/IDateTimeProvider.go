package Interfaces

import "time"

type IDateTimeProvider interface {
	UTCNow() time.Time
}



