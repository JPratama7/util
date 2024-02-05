package sets

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddAndContains(t *testing.T) {
	set := NewSets[int]()
	set.Add(1)
	assert.True(t, set.Contains(1), "Expected set to contain 1")
}

func TestRemove(t *testing.T) {
	set := NewSets[int]()
	set.Add(1)
	set.Remove(1)
	assert.False(t, set.Contains(1), "Expected set to not contain 1")
}

func TestSize(t *testing.T) {
	set := NewSets[int]()
	set.Add(1)
	set.Add(2)
	assert.Equal(t, 2, set.Size(), "Expected set size to be 2")
}

func TestClear(t *testing.T) {
	set := NewSets[int]()
	set.Add(1)
	set.Clear()
	assert.Equal(t, 0, set.Size(), "Expected set size to be 0 after clear")
}

func TestUnion(t *testing.T) {
	set1 := NewSets[int]()
	set1.Add(1)
	set1.Add(2)
	set2 := NewSets[int]()
	set2.Add(2)
	set2.Add(3)
	union := set1.Union(set2)
	assert.Equal(t, 3, union.Size(), "Expected union size to be 3")
	assert.True(t, union.Contains(1), "Expected union to contain 1")
	assert.True(t, union.Contains(2), "Expected union to contain 2")
	assert.True(t, union.Contains(3), "Expected union to contain 3")
}

func TestIntersection(t *testing.T) {
	set1 := NewSets[int]()
	set1.Add(1)
	set1.Add(2)
	set2 := NewSets[int]()
	set2.Add(2)
	set2.Add(3)
	intersection := set1.Intersection(set2)
	assert.Equal(t, 1, intersection.Size(), "Expected intersection size to be 1")
	assert.True(t, intersection.Contains(2), "Expected intersection to contain 2")
}

func TestDifference(t *testing.T) {
	set1 := NewSets[int]()
	set1.Add(1)
	set1.Add(2)
	set2 := NewSets[int]()
	set2.Add(2)
	set2.Add(3)
	difference := set1.Difference(set2)
	assert.Equal(t, 1, difference.Size(), "Expected difference size to be 1")
	assert.True(t, difference.Contains(1), "Expected difference to contain 1")
}

func TestIsSubset(t *testing.T) {
	set1 := NewSets[int]()
	set1.Add(1)
	set1.Add(2)
	set2 := NewSets[int]()
	set2.Add(1)
	set2.Add(2)
	set2.Add(3)
	assert.True(t, set1.IsSubset(set2), "Expected set1 to be a subset of set2")
}

func TestIter(t *testing.T) {
	set := NewSets[int]()
	set.Add(1)
	set.Add(2)
	ch := set.Next()
	val1 := <-ch
	val2 := <-ch
	assert.Contains(t, []int{1, 2}, val1, "Expected to receive 1 or 2 from iterator")
	assert.Contains(t, []int{1, 2}, val2, "Expected to receive 1 or 2 from iterator")
	assert.NotEqual(t, val1, val2, "Expected to receive different values from iterator")
}
