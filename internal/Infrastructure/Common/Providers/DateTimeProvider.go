package Providers

import (
	"time"

	. "github.com/stellayazilim/neptune_cms/internal/Application/Common/Interfaces"
)

type dateTimeProvider time.Time

func (d *dateTimeProvider) UTCNow() time.Time {
	return time.Now()
}

func DateTimeProvider() IDateTimeProvider {
	return &dateTimeProvider{}
}
