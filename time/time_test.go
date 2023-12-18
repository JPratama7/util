package time

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNowReturnsCurrentTime(t *testing.T) {
	result := time.Now()
	assert.WithinDuration(t, time.Now(), result, 1*time.Second)
}

func TestAddCenturiesAddsCorrectly(t *testing.T) {
	now := time.Now()
	result := AddCenturies(now, 1)
	expected := now.AddDate(100, 0, 0)
	assert.Equal(t, expected, result)
}

func TestAddCenturyAddsCorrectly(t *testing.T) {
	now := time.Now()
	result := AddCentury(now)
	expected := now.AddDate(100, 0, 0)
	assert.Equal(t, expected, result)
}

func TestAddDayAddsCorrectly(t *testing.T) {
	now := time.Now()
	result := AddDay(now)
	expected := now.AddDate(0, 0, 1)
	assert.Equal(t, expected, result)
}

func TestAddDaysAddsCorrectly(t *testing.T) {
	now := time.Now()
	result := AddDays(now, 5)
	expected := now.AddDate(0, 0, 5)
	assert.Equal(t, expected, result)
}

func TestAddDecadeAddsCorrectly(t *testing.T) {
	now := time.Now()
	result := AddDecade(now)
	expected := now.AddDate(10, 0, 0)
	assert.Equal(t, expected, result)
}

func TestAddDecadesAddsCorrectly(t *testing.T) {
	now := time.Now()
	result := AddDecades(now, 3)
	expected := now.AddDate(30, 0, 0)
	assert.Equal(t, expected, result)
}

func TestAddHourAddsCorrectly(t *testing.T) {
	now := time.Now()
	result := AddHour(now)
	expected := now.Add(time.Hour)
	assert.Equal(t, expected, result)
}

func TestAddHoursAddsCorrectly(t *testing.T) {
	now := time.Now()
	result := AddHours(now, 3)
	expected := now.Add(3 * time.Hour)
	assert.Equal(t, expected, result)
}
