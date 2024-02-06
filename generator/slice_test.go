package generator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTestNewIterSliceInitialization(t *testing.T) {
	iter := []int{1, 2, 3}
	iterSlice := NewIterSlice(iter)

	assert.NotNil(t, iterSlice)
}

func TestNextWithEmptySlice(t *testing.T) {
	iter := make([]int, 0)
	iterSlice := NewIterSlice(iter)

	hasNext := iterSlice.Next()

	assert.False(t, hasNext)
}

func TestNextWithNonEmptySlice(t *testing.T) {
	iter := []int{1, 2, 3}
	iterSlice := NewIterSlice(iter)

	hasNext := iterSlice.Next()

	assert.True(t, hasNext)
}

func TestNextUntilExhaustionSlice(t *testing.T) {
	iter := []int{1, 2, 3}
	iterSlice := NewIterSlice(iter)

	for i := 0; i < len(iter); i++ {
		assert.True(t, iterSlice.Next())
	}

	assert.False(t, iterSlice.Next())
}

func TestValueWithEmptySlice(t *testing.T) {
	iter := make([]int, 0)
	iterSlice := NewIterSlice(iter)

	value := iterSlice.Value()

	assert.Empty(t, value)
}

func TestValueWithNonEmptySlice(t *testing.T) {
	iter := []int{1, 2, 3}
	iterSlice := NewIterSlice(iter)

	iterSlice.Next()
	value := iterSlice.Value()

	assert.NotNil(t, value)
	assert.Equal(t, 1, value)
}

func TestHasNextWithEmptySlice(t *testing.T) {
	iter := make([]int, 0)
	iterSlice := NewIterSlice(iter)

	hasNext := iterSlice.HasNext()

	assert.False(t, hasNext)
}

func TestHasNextWithNonEmptySlice(t *testing.T) {
	iter := []int{1, 2, 3}
	iterSlice := NewIterSlice(iter)

	hasNext := iterSlice.HasNext()

	assert.True(t, hasNext)
}

func TestHasNextUntilExhaustion(t *testing.T) {
	iter := []int{1, 2, 3}
	iterSlice := NewIterSlice(iter)

	for i := 0; i < len(iter); i++ {
		assert.True(t, iterSlice.HasNext())
		iterSlice.Next()
	}

	assert.False(t, iterSlice.HasNext())
}
