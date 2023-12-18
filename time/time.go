package time

import (
	"github.com/golang-module/carbon"
	"time"
)

func Now() time.Time {
	return time.Now()
}

func AddCenturies(t time.Time, n int) (v time.Time) {
	return carbon.FromStdTime(t).AddCenturies(n).ToStdTime()
}

func AddCentury(t time.Time) (v time.Time) {
	return carbon.FromStdTime(t).AddCentury().ToStdTime()
}

func AddDay(t time.Time) (v time.Time) {
	return carbon.FromStdTime(t).AddDay().ToStdTime()
}

func AddDays(t time.Time, n int) (v time.Time) {
	return carbon.FromStdTime(t).AddDays(n).ToStdTime()
}

func AddDecade(t time.Time) (v time.Time) {
	return carbon.FromStdTime(t).AddDecade().ToStdTime()
}

func AddDecades(t time.Time, n int) (v time.Time) {
	return carbon.FromStdTime(t).AddDecades(n).ToStdTime()
}

func AddHour(t time.Time) (v time.Time) {
	return carbon.FromStdTime(t).AddHour().ToStdTime()
}

func AddHours(t time.Time, n int) (v time.Time) {
	return carbon.FromStdTime(t).AddHours(n).ToStdTime()
}
