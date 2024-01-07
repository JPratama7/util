package time

import (
	"github.com/golang-module/carbon/v2"
	"time"
)

func Now() time.Time {
	return time.Now()
}

func AddCenturies(t time.Time, n int) (v time.Time) {
	return carbon.CreateFromStdTime(t).AddCenturies(n).ToStdTime()
}

func AddCentury(t time.Time) (v time.Time) {
	return carbon.CreateFromStdTime(t).AddCentury().ToStdTime()
}

func AddDay(t time.Time) (v time.Time) {
	return carbon.CreateFromStdTime(t).AddDay().ToStdTime()
}

func AddDays(t time.Time, n int) (v time.Time) {
	return carbon.CreateFromStdTime(t).AddDays(n).ToStdTime()
}

func AddDecade(t time.Time) (v time.Time) {
	return carbon.CreateFromStdTime(t).AddDecade().ToStdTime()
}

func AddDecades(t time.Time, n int) (v time.Time) {
	return carbon.CreateFromStdTime(t).AddDecades(n).ToStdTime()
}

func AddHour(t time.Time) (v time.Time) {
	return carbon.CreateFromStdTime(t).AddHour().ToStdTime()
}

func AddHours(t time.Time, n int) (v time.Time) {
	return carbon.CreateFromStdTime(t).AddHours(n).ToStdTime()
}

func ConvertToDuration(t time.Time, round ...time.Duration) (v time.Duration) {
	rounder := time.Second
	if len(round) > 0 {
		rounder = round[0]
	}

	return time.Until(t).Round(rounder)
}
