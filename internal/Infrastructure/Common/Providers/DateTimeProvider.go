package Providers

import (
	"time"

	"github.com/stellayazilim/neptune.application/Common/Interfaces"
)

type dateTimeProvider time.Time

func (d *dateTimeProvider) UTCNow() time.Time {
	return time.Now()
}

func DateTimeProvider() Interfaces.IDateTimeProvider {
	return &dateTimeProvider{}
}
