package sets

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func UnsafeSetsOperations(t *testing.T) {
	t.Parallel()

	t.Run("ShouldAddAndContainElement", func(t *testing.T) {
		set := UnsafeUnsafeSets[int]()
		set.Add(1)
		assert.True(t, set.Contains(1))
	})

	t.Run("ShouldRemoveElement", func(t *testing.T) {
		set := UnsafeUnsafeSets[int]()
		set.Add(1)
		set.Remove(1)
		assert.False(t, set.Contains(1))
	})

	t.Run("ShouldReturnCorrectSize", func(t *testing.T) {
		set := UnsafeUnsafeSets[int]()
		set.Add(1)
		set.Add(2)
		assert.Equal(t, 2, set.Size())
	})

	t.Run("ShouldClearSet", func(t *testing.T) {
		set := UnsafeUnsafeSets[int]()
		set.Add(1)
		set.Clear()
		assert.Equal(t, 0, set.Size())
	})

	t.Run("ShouldReturnUnionOfSets", func(t *testing.T) {
		set1 := UnsafeUnsafeSets[int]()
		set1.Add(1)
		set1.Add(2)
		set2 := UnsafeUnsafeSets[int]()
		set2.Add(2)
		set2.Add(3)
		union := set1.Union(set2)
		assert.Equal(t, 3, union.Size())
		assert.True(t, union.Contains(1))
		assert.True(t, union.Contains(2))
		assert.True(t, union.Contains(3))
	})

	t.Run("ShouldReturnIntersectionOfSets", func(t *testing.T) {
		set1 := UnsafeUnsafeSets[int]()
		set1.Add(1)
		set1.Add(2)
		set2 := UnsafeUnsafeSets[int]()
		set2.Add(2)
		set2.Add(3)
		intersection := set1.Intersection(set2)
		assert.Equal(t, 1, intersection.Size())
		assert.True(t, intersection.Contains(2))
	})

	t.Run("ShouldReturnDifferenceOfSets", func(t *testing.T) {
		set1 := UnsafeUnsafeSets[int]()
		set1.Add(1)
		set1.Add(2)
		set2 := UnsafeUnsafeSets[int]()
		set2.Add(2)
		set2.Add(3)
		difference := set1.Difference(set2)
		assert.Equal(t, 1, difference.Size())
		assert.True(t, difference.Contains(1))
	})

	t.Run("ShouldCheckIfSetIsSubset", func(t *testing.T) {
		set1 := UnsafeUnsafeSets[int]()
		set1.Add(1)
		set1.Add(2)
		set2 := UnsafeUnsafeSets[int]()
		set2.Add(1)
		set2.Add(2)
		set2.Add(3)
		assert.True(t, set1.IsSubset(set2))
	})

	t.Run("ShouldIterateOverSet", func(t *testing.T) {
		set := UnsafeUnsafeSets[int]()
		set.Add(1)
		set.Add(2)
		ch := set.Next()
		val1 := <-ch
		val2 := <-ch
		assert.Contains(t, []int{1, 2}, val1)
		assert.Contains(t, []int{1, 2}, val2)
		assert.NotEqual(t, val1, val2)
	})
}
