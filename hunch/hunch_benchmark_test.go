package hunch

import (
	"context"
	"testing"
)

func BenchmarkTakeWithOnlyOne(b *testing.B) {
	rootCtx := context.Background()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Take(
			rootCtx,
			1,
			func(ctx context.Context) (int, error) {
				return 1, nil
			},
		)
	}
}

func BenchmarkTakeWithOnlyOneMut(b *testing.B) {
	rootCtx := context.Background()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TakeMut(
			rootCtx,
			1,
			func(ctx context.Context) (int, error) {
				return 1, nil
			},
		)
	}
}

func BenchmarkTakeWithFiveExecsThatNeedsOne(b *testing.B) {
	rootCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Take(
			rootCtx,
			1,
			func(ctx context.Context) (int, error) {
				return 1, nil
			},
			func(ctx context.Context) (int, error) {
				return 2, nil
			},
			func(ctx context.Context) (int, error) {
				return 3, nil
			},
			func(ctx context.Context) (int, error) {
				return 4, nil
			},
			func(ctx context.Context) (int, error) {
				return 5, nil
			},
		)
	}
}
func BenchmarkTakeWithFiveExecsThatNeedsOneMut(b *testing.B) {
	rootCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TakeMut(
			rootCtx,
			1,
			func(ctx context.Context) (int, error) {
				return 1, nil
			},
			func(ctx context.Context) (int, error) {
				return 2, nil
			},
			func(ctx context.Context) (int, error) {
				return 3, nil
			},
			func(ctx context.Context) (int, error) {
				return 4, nil
			},
			func(ctx context.Context) (int, error) {
				return 5, nil
			},
		)
	}
}

func BenchmarkTakeWithFiveExecsThatNeedsFive(b *testing.B) {
	rootCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Take(
			rootCtx,
			5,
			func(ctx context.Context) (int, error) {
				return 1, nil
			},
			func(ctx context.Context) (int, error) {
				return 2, nil
			},
			func(ctx context.Context) (int, error) {
				return 3, nil
			},
			func(ctx context.Context) (int, error) {
				return 4, nil
			},
			func(ctx context.Context) (int, error) {
				return 5, nil
			},
		)
	}
}

func BenchmarkTakeWithFiveExecsThatNeedsFiveMut(b *testing.B) {
	rootCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TakeMut(
			rootCtx,
			5,
			func(ctx context.Context) (int, error) {
				return 1, nil
			},
			func(ctx context.Context) (int, error) {
				return 2, nil
			},
			func(ctx context.Context) (int, error) {
				return 3, nil
			},
			func(ctx context.Context) (int, error) {
				return 4, nil
			},
			func(ctx context.Context) (int, error) {
				return 5, nil
			},
		)
	}
}

func BenchmarkAllWithFiveExecs(b *testing.B) {
	rootCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		All(
			rootCtx,
			func(ctx context.Context) (int, error) {
				return 1, nil
			},
			func(ctx context.Context) (int, error) {
				return 2, nil
			},
			func(ctx context.Context) (int, error) {
				return 3, nil
			},
			func(ctx context.Context) (int, error) {
				return 4, nil
			},
			func(ctx context.Context) (int, error) {
				return 5, nil
			},
		)
	}
}
func BenchmarkAllWithFiveExecsMut(b *testing.B) {
	rootCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		AllMut(
			rootCtx,
			func(ctx context.Context) (int, error) {
				return 1, nil
			},
			func(ctx context.Context) (int, error) {
				return 2, nil
			},
			func(ctx context.Context) (int, error) {
				return 3, nil
			},
			func(ctx context.Context) (int, error) {
				return 4, nil
			},
			func(ctx context.Context) (int, error) {
				return 5, nil
			},
		)
	}
}

func BenchmarkThrow(b *testing.B) {
	// Create a slice of Executable functions for testing
	execs := make([]Executable[int], 0, b.N)
	for i := range execs {
		execs = append(execs, func(ctx context.Context) (int, error) {
			return i, nil
		})
	}

	// Run the Throw function b.N times
	for i := 0; i < b.N; i++ {
		Throw(context.Background(), execs...)
	}
}
func BenchmarkThrowMut(b *testing.B) {
	// Create a slice of Executable functions for testing
	execs := make([]Executable[int], 0, 5000)
	for i := range execs {
		execs = append(execs, func(ctx context.Context) (int, error) {
			return i, nil
		})
	}

	// Run the Throw function b.N times
	for i := 0; i < b.N; i++ {
		ThrowMut(context.Background(), execs...)
	}
}
