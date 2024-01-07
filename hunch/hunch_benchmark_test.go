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

func BenchmarkAllWithFiveExecs(b *testing.B) {
	rootCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		All[int](
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
