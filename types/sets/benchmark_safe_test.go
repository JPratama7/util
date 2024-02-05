package sets

import "testing"

func BenchmarkSetsAdd(b *testing.B) {
	set := NewSets[int]()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set.Add(i)
	}
}

func BenchmarkSetsRemove(b *testing.B) {
	set := NewSets[int]()
	for i := 0; i < b.N; i++ {
		set.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set.Remove(i)
	}
}

func BenchmarkSetsContains(b *testing.B) {
	set := NewSets[int]()
	for i := 0; i < b.N; i++ {
		set.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set.Contains(i)
	}
}

func BenchmarkSetsSize(b *testing.B) {
	set := NewSets[int]()
	for i := 0; i < b.N; i++ {
		set.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set.Size()
	}
}

func BenchmarkSetsToSlice(b *testing.B) {
	set := NewSets[int]()
	for i := 0; i < b.N; i++ {
		set.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set.ToSlice()
	}
}

func BenchmarkSetsClear(b *testing.B) {
	set := NewSets[int]()
	for i := 0; i < b.N; i++ {
		set.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set.Clear()
	}
}

func BenchmarkSetsUnion(b *testing.B) {
	set1 := NewSets[int]()
	set2 := NewSets[int]()
	for i := 0; i < b.N; i++ {
		set1.Add(i)
		set2.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set1.Union(set2)
	}
}

func BenchmarkSetsIntersection(b *testing.B) {
	set1 := NewSets[int]()
	set2 := NewSets[int]()
	for i := 0; i < b.N; i++ {
		set1.Add(i)
		set2.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set1.Intersection(set2)
	}
}

func BenchmarkSetsDifference(b *testing.B) {
	set1 := NewSets[int]()
	set2 := NewSets[int]()
	for i := 0; i < b.N; i++ {
		set1.Add(i)
		set2.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set1.Difference(set2)
	}
}

func BenchmarkSetsIsSubset(b *testing.B) {
	set1 := NewSets[int]()
	set2 := NewSets[int]()
	for i := 0; i < b.N; i++ {
		set1.Add(i)
		set2.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set1.IsSubset(set2)
	}
}

func BenchmarkSetsNext(b *testing.B) {
	set := NewSets[int]()
	for i := 0; i < b.N; i++ {
		set.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set.Next()
	}
}

func BenchmarkSetsRange(b *testing.B) {
	set := NewSets[int]()
	for i := 0; i < b.N; i++ {
		set.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set.Range(func(v int) bool {
			return true
		})
	}
}
