package hunch

import (
	"context"
	"testing"
	"time"
)

func BenchmarkFirst(b *testing.B) {
	ctx := context.Background()
	exec1 := mockExecutable{value: 1, delay: 1 * time.Second}
	exec2 := mockExecutable{value: 2, delay: 2 * time.Second}

	for i := 0; i < b.N; i++ {
		_, _ = First(ctx, exec1.Execute, exec2.Execute)
	}
}

func BenchmarkThrow(b *testing.B) {
	ctx := context.Background()
	exec1 := mockExecutable{value: 1, delay: 1 * time.Second}
	exec2 := mockExecutable{value: 2, delay: 2 * time.Second}

	for i := 0; i < b.N; i++ {
		Throw(ctx, exec1.Execute, exec2.Execute)
	}
}

func BenchmarkTake(b *testing.B) {
	ctx := context.Background()
	exec1 := mockExecutable{value: 1, delay: 1 * time.Second}
	exec2 := mockExecutable{value: 2, delay: 2 * time.Second}

	for i := 0; i < b.N; i++ {
		_, _ = Take(ctx, 2, exec1.Execute, exec2.Execute)
	}
}

func BenchmarkAll(b *testing.B) {
	ctx := context.Background()
	exec1 := mockExecutable{value: 1, delay: 1 * time.Second}
	exec2 := mockExecutable{value: 2, delay: 2 * time.Second}
	exec3 := mockExecutable{value: 3, delay: 3 * time.Second}

	for i := 0; i < b.N; i++ {
		_, _ = All(ctx, exec1.Execute, exec2.Execute, exec3.Execute)
	}
}
