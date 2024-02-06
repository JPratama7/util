package generator

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestNewIterMap(t *testing.T) {
	iter := map[int]string{1: "one", 2: "two", 3: "three"}
	iterMap := NewIterMap(iter)

	assert.NotNil(t, iterMap)
	assert.Equal(t, iter, iterMap.iter)
	assert.EqualValues(t, 0, iterMap.lastKey)
}

func TestNextWithEmptyMap(t *testing.T) {
	iter := make(map[int]string)
	iterMap := NewIterMap(iter)

	hasNext := iterMap.Next()

	assert.False(t, hasNext)
}

func TestNextWithNonEmptyMap(t *testing.T) {
	iter := map[int]string{1: "one", 2: "two", 3: "three"}
	iterMap := NewIterMap(iter)

	hasNext := iterMap.Next()

	assert.True(t, hasNext)
}

func TestNextUntilExhaustion(t *testing.T) {
	iter := map[int]string{1: "one", 2: "two", 3: "three"}
	iterMap := NewIterMap(iter)

	assert.True(t, iterMap.Next())
	for i := 0; iterMap.HasNext(); i++ {
		assert.True(t, iterMap.Next())
	}

	assert.False(t, iterMap.Next())
}

func TestValueWithEmptyMap(t *testing.T) {
	iter := make(map[int]string)
	iterMap := NewIterMap(iter)

	key, value := iterMap.Value()

	assert.Empty(t, key)
	assert.Empty(t, value)
}

func TestValueWithNonEmptyMap(t *testing.T) {
	iter := map[int]string{1: "one", 2: "two", 3: "three"}
	iterMap := NewIterMap(iter)

	iterMap.Next()
	key, value := iterMap.Value()
	log.Printf("key: %v, value: %v", key, value)

	assert.NotNil(t, key)
	assert.NotNil(t, value)
}
