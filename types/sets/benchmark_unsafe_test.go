package sets

import "testing"

func BenchmarkUnsafeSets(b *testing.B) {
	set := UnsafeUnsafeSets[int]()
	for i := 0; i < b.N; i++ {
		set.Add(i)
	}
}

func BenchmarkUnsafeSetsRemove(b *testing.B) {
	set := UnsafeUnsafeSets[int]()
	for i := 0; i < b.N; i++ {
		set.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set.Remove(i)
	}
}

func BenchmarkUnsafeSetsContains(b *testing.B) {
	set := UnsafeUnsafeSets[int]()
	for i := 0; i < b.N; i++ {
		set.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set.Contains(i)
	}
}

func BenchmarkUnsafeSetsSize(b *testing.B) {
	set := UnsafeUnsafeSets[int]()
	for i := 0; i < b.N; i++ {
		set.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set.Size()
	}
}

func BenchmarkUnsafeSetsToSlice(b *testing.B) {
	set := UnsafeUnsafeSets[int]()
	for i := 0; i < b.N; i++ {
		set.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set.ToSlice()
	}
}

func BenchmarkUnsafeSetsClear(b *testing.B) {
	set := UnsafeUnsafeSets[int]()
	for i := 0; i < b.N; i++ {
		set.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set.Clear()
	}
}

func BenchmarkUnsafeSetsUnion(b *testing.B) {
	set1 := UnsafeUnsafeSets[int]()
	set2 := UnsafeUnsafeSets[int]()
	for i := 0; i < b.N; i++ {
		set1.Add(i)
		set2.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set1.Union(set2)
	}
}

func BenchmarkUnsafeSetsIntersection(b *testing.B) {
	set1 := UnsafeUnsafeSets[int]()
	set2 := UnsafeUnsafeSets[int]()
	for i := 0; i < b.N; i++ {
		set1.Add(i)
		set2.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set1.Intersection(set2)
	}
}

func BenchmarkUnsafeSetsDifference(b *testing.B) {
	set1 := UnsafeUnsafeSets[int]()
	set2 := UnsafeUnsafeSets[int]()
	for i := 0; i < b.N; i++ {
		set1.Add(i)
		set2.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set1.Difference(set2)
	}
}

func BenchmarkUnsafeSetsIsSubset(b *testing.B) {
	set1 := UnsafeUnsafeSets[int]()
	set2 := UnsafeUnsafeSets[int]()
	for i := 0; i < b.N; i++ {
		set1.Add(i)
		set2.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set1.IsSubset(set2)
	}
}

func BenchmarkUnsafeSetsNext(b *testing.B) {
	set := UnsafeUnsafeSets[int]()
	for i := 0; i < b.N; i++ {
		set.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set.Next()
	}
}
