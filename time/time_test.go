package time

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNowReturnsCurrentTime(t *testing.T) {
	result := Now()
	assert.WithinDuration(t, time.Now(), result, 1*time.Second)
}

func TestAddCenturiesAddsCorrectly(t *testing.T) {
	now := time.Now()
	result := AddCenturies(now, 1)
	expected := now.AddDate(100, 0, 0)
	assert.WithinDuration(t, expected, result, 1*time.Second)
}

func TestAddCenturyAddsCorrectly(t *testing.T) {
	now := time.Now()
	result := AddCentury(now)
	expected := now.AddDate(100, 0, 0)
	assert.WithinDuration(t, expected, result, 1*time.Second)
}

func TestAddDayAddsCorrectly(t *testing.T) {
	now := time.Now()
	result := AddDay(now)
	expected := now.AddDate(0, 0, 1)
	assert.WithinDuration(t, expected, result, 1*time.Second)
}

func TestAddDaysAddsCorrectly(t *testing.T) {
	now := time.Now()
	result := AddDays(now, 5)
	expected := now.AddDate(0, 0, 5)
	assert.WithinDuration(t, expected, result, 1*time.Second)
}

func TestAddDecadeAddsCorrectly(t *testing.T) {
	now := time.Now()
	result := AddDecade(now)
	expected := now.AddDate(10, 0, 0)
	assert.WithinDuration(t, expected, result, 1*time.Second)
}

func TestAddDecadesAddsCorrectly(t *testing.T) {
	now := time.Now()
	result := AddDecades(now, 3)
	expected := now.AddDate(30, 0, 0)
	assert.WithinDuration(t, expected, result, 1*time.Second)
}

func TestAddHourAddsCorrectly(t *testing.T) {
	now := time.Now()
	result := AddHour(now)
	expected := time.Now().Add(time.Hour)
	assert.WithinDuration(t, expected, result, 1*time.Second)
}

func TestAddHoursAddsCorrectly(t *testing.T) {
	now := time.Now()
	result := AddHours(now, 3)
	expected := now.Add(3 * time.Hour)
	assert.WithinDuration(t, expected, result, 1*time.Second)
}

func TestConvertToDurationReturnsCorrectDurationForFutureTime(t *testing.T) {
	futureTime := time.Now().Add(3 * time.Hour)
	result := ConvertToDuration(futureTime)
	expected := 3 * time.Hour
	assert.Equal(t, expected, result)
}

func TestConvertToDurationReturnsCorrectDurationForPastTime(t *testing.T) {
	pastTime := time.Now().Add(-3 * time.Hour)
	result := ConvertToDuration(pastTime)
	expected := -3 * time.Hour
	assert.Equal(t, expected, result)
}

func TestConvertToDurationReturnsZeroForCurrentTime(t *testing.T) {
	currentTime := time.Now()
	result := ConvertToDuration(currentTime)
	expected := 0 * time.Hour
	assert.Equal(t, expected, result)
}
