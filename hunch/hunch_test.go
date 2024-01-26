package hunch

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"sync/atomic"
	"testing"
	"time"
)

type AppError struct {
}

func (err AppError) Error() string {
	return "app error!"
}

type MultiReturns struct {
	Val interface{}
	Err error
}

func isSlice(v interface{}) bool {
	rv := reflect.ValueOf(v)
	return rv.Kind() == reflect.Slice
}

func TestTake_ShouldWorksAsExpected(t *testing.T) {
	t.Parallel()

	rootCtx := context.Background()
	ch := make(chan MultiReturns)
	go func() {
		r, err := Take(rootCtx, 3, func(ctx context.Context) (interface{}, error) {
			time.Sleep(200 * time.Millisecond)
			return 1, nil
		}, func(ctx context.Context) (interface{}, error) {
			time.Sleep(100 * time.Millisecond)
			return 2, nil
		}, func(ctx context.Context) (interface{}, error) {
			<-time.After(300 * time.Millisecond)
			return 3, nil
		})

		ch <- MultiReturns{r, err}
		close(ch)
	}()

	r := <-ch
	if r.Err != nil {
		t.Errorf("Gets an error: %v\n", r.Err)
	}

	rs := []int{}
	for _, v := range r.Val.([]interface{}) {
		rs = append(rs, v.(int))
	}

	assert.Equal(t, []int{2, 1, 3}, rs)
}

func TestTake_ShouldJustTakeAll(t *testing.T) {
	t.Parallel()

	rootCtx := context.Background()
	ch := make(chan MultiReturns)
	go func() {
		r, err := Take(rootCtx, 100, func(ctx context.Context) (interface{}, error) {
			return 1, nil
		}, func(ctx context.Context) (interface{}, error) {
			return 2, nil
		}, func(ctx context.Context) (interface{}, error) {
			return 3, nil
		})

		ch <- MultiReturns{r, err}
		close(ch)
	}()

	r := <-ch
	if r.Err != nil {
		t.Errorf("Gets an error: %v\n", r.Err)
	}

	rs := []int{}
	for _, v := range r.Val.([]interface{}) {
		rs = append(rs, v.(int))
	}

	assert.Equal(t, 3, len(rs))
}

func TestTake_ShouldLimitResults(t *testing.T) {
	t.Parallel()

	rootCtx := context.Background()
	ch := make(chan MultiReturns)
	go func() {
		r, err := Take(rootCtx, 2, func(ctx context.Context) (interface{}, error) {
			time.Sleep(200 * time.Millisecond)
			return 1, nil
		}, func(ctx context.Context) (interface{}, error) {
			time.Sleep(100 * time.Millisecond)
			return 2, nil
		}, func(ctx context.Context) (interface{}, error) {
			<-time.After(300 * time.Millisecond)
			return 3, nil
		})

		ch <- MultiReturns{r, err}
		close(ch)
	}()

	r := <-ch
	rs := []int{}
	for _, v := range r.Val.([]interface{}) {
		rs = append(rs, v.(int))
	}

	assert.Equal(t, []int{2, 1}, rs)
}

func TestTake_ShouldCancelWhenOneExecutableReturnedError(t *testing.T) {
	t.Parallel()

	rootCtx := context.Background()
	ch := make(chan MultiReturns)
	go func() {
		r, err := Take(rootCtx, 3, func(ctx context.Context) (int, error) {
			time.Sleep(100 * time.Millisecond)
			return 1, nil
		}, func(ctx context.Context) (int, error) {
			time.Sleep(200 * time.Millisecond)
			return 0, AppError{}
		}, func(ctx context.Context) (int, error) {
			time.Sleep(200 * time.Millisecond)
			return 3, nil
		})

		ch <- MultiReturns{r, err}
		close(ch)
	}()

	r := <-ch
	assert.False(t, isSlice(r.Val) && len(r.Val.([]int)) != 0, "Return Value should be default, gets: \"%v\"\n", r.Val)
	assert.NotNil(t, r.Err, "Should returns an Error, gets `nil`\n")
}

func TestTake_ShouldCancelWhenRootCanceled(t *testing.T) {
	t.Parallel()

	rootCtx, cancel := context.WithCancel(context.Background())
	ch := make(chan MultiReturns)
	go func() {
		r, err := Take(rootCtx, 3, func(ctx context.Context) (int, error) {
			time.Sleep(100 * time.Millisecond)
			return 1, nil
		}, func(ctx context.Context) (int, error) {
			time.Sleep(200 * time.Millisecond)
			return 2, nil
		}, func(ctx context.Context) (int, error) {
			select {
			case <-ctx.Done():
				return 0, AppError{}
			case <-time.After(300 * time.Millisecond):
				return 3, nil
			}
		})

		ch <- MultiReturns{r, err}
		close(ch)
	}()

	go func() {
		time.Sleep(150 * time.Millisecond)
		cancel()
	}()

	r := <-ch
	assert.False(t, isSlice(r.Val) && len(r.Val.([]int)) != 0, "Return Value should be default, gets: \"%v\"\n", r.Val)
	assert.NotNil(t, r.Err, "Should returns an Error, gets `nil`\n")
}

func TestAll_ShouldWorksAsExpected(t *testing.T) {
	t.Parallel()

	rootCtx := context.Background()
	ch := make(chan MultiReturns)
	go func() {
		r, err := All(rootCtx, func(ctx context.Context) (int, error) {
			time.Sleep(200 * time.Millisecond)
			return 1, nil
		}, func(ctx context.Context) (int, error) {
			time.Sleep(100 * time.Millisecond)
			return 2, nil
		}, func(ctx context.Context) (int, error) {
			<-time.After(300 * time.Millisecond)
			return 3, nil
		})

		ch <- MultiReturns{r, err}
		close(ch)
	}()

	r := <-ch
	if r.Err != nil {
		t.Errorf("Gets an error: %v\n", r.Err)
	}

	var rs []int
	rs = append(rs, r.Val.([]int)...)

	assert.Equal(t, []int{1, 2, 3}, rs)
}

func TestAll_WhenOutOfOrder(t *testing.T) {
	t.Parallel()

	rootCtx := context.Background()
	ch := make(chan MultiReturns)
	go func() {
		r, err := All(rootCtx, func(ctx context.Context) (int, error) {
			time.Sleep(200 * time.Millisecond)
			return 1, nil
		}, func(ctx context.Context) (int, error) {
			time.Sleep(300 * time.Millisecond)
			return 2, nil
		}, func(ctx context.Context) (int, error) {
			<-time.After(100 * time.Millisecond)
			return 3, nil
		})

		ch <- MultiReturns{r, err}
		close(ch)
	}()

	r := <-ch
	if r.Err != nil {
		t.Errorf("Gets an error: %v\n", r.Err)
	}

	var rs []int
	rs = append(rs, r.Val.([]int)...)

	assert.Equal(t, []int{1, 2, 3}, rs)
}

func TestAll_WhenRootCtxCanceled(t *testing.T) {
	t.Parallel()

	rootCtx, cancel := context.WithCancel(context.Background())
	ch := make(chan MultiReturns)
	go func() {
		r, err := All(rootCtx, func(ctx context.Context) (interface{}, error) {
			time.Sleep(100 * time.Millisecond)
			return 1, nil
		}, func(ctx context.Context) (interface{}, error) {
			time.Sleep(200 * time.Millisecond)
			return 2, nil
		}, func(ctx context.Context) (interface{}, error) {
			select {
			case <-ctx.Done():
				return nil, AppError{}
			case <-time.After(300 * time.Millisecond):
				return 3, nil
			}
		})

		ch <- MultiReturns{r, err}
		close(ch)
	}()

	go func() {
		time.Sleep(150 * time.Millisecond)
		cancel()
	}()

	r := <-ch
	assert.True(t, isSlice(r.Val), "Return Value should be default, gets: \"%v\"\n", r.Val)
	assert.Empty(t, r.Val.([]interface{}), "Return Value should be empty")
	assert.NotNil(t, r.Err, "Should returns an Error, gets `nil`\n")
}

func TestAll_WhenAnySubFunctionFailed(t *testing.T) {
	t.Parallel()

	rootCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch := make(chan MultiReturns)
	go func() {
		r, err := All(rootCtx, func(ctx context.Context) (interface{}, error) {
			time.Sleep(100 * time.Millisecond)
			return 1, nil
		}, func(ctx context.Context) (interface{}, error) {
			time.Sleep(200 * time.Millisecond)
			return 2, nil
		}, func(ctx context.Context) (interface{}, error) {
			time.Sleep(300 * time.Millisecond)
			return nil, AppError{}
		})

		ch <- MultiReturns{r, err}
		close(ch)
	}()

	r := <-ch
	assert.True(t, isSlice(r.Val), "Return Value should be default, gets: \"%v\"\n", r.Val)
	assert.Empty(t, r.Val.([]interface{}), "Return Value should be empty")
	assert.NotNil(t, r.Err, "Should returns an Error, gets `nil`\n")
	assert.Equal(t, "app error!", r.Err.Error(), "Should returns an AppError, gets \"%v\"\n", r.Err.Error())
}

func TestRetry_WithNoFailure(t *testing.T) {
	t.Parallel()

	times := 0
	expect := 1

	rootCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch := make(chan MultiReturns)
	go func() {
		r, err := Retry(rootCtx, 10, func(ctx context.Context) (interface{}, error) {
			times++

			time.Sleep(100 * time.Millisecond)
			return expect, nil
		})

		ch <- MultiReturns{r, err}
		close(ch)
	}()

	r := <-ch
	assert.Nil(t, r.Err)

	assert.IsTypef(t, 1, r.Val, "Should return int, but got: %v", r.Val)

	assert.Equal(t, 1, times)
}

func TestLast_ShouldWorksAsExpected(t *testing.T) {
	t.Parallel()

	rootCtx := context.Background()
	ch := make(chan MultiReturns)
	go func() {
		r, err := Last(rootCtx, 2, func(ctx context.Context) (interface{}, error) {
			time.Sleep(200 * time.Millisecond)
			return 1, nil
		}, func(ctx context.Context) (interface{}, error) {
			time.Sleep(100 * time.Millisecond)
			return 2, nil
		}, func(ctx context.Context) (interface{}, error) {
			<-time.After(300 * time.Millisecond)
			return 3, nil
		})

		ch <- MultiReturns{r, err}
		close(ch)
	}()

	r := <-ch
	if r.Err != nil {
		t.Errorf("Gets an error: %v\n", r.Err)
	}

	rs := []int{}
	for _, v := range r.Val.([]interface{}) {
		rs = append(rs, v.(int))
	}

	assert.Equal(t, []int{1, 3}, rs)

}

func TestWaterfall_ShouldWorksAsExpected(t *testing.T) {
	t.Parallel()

	rootCtx := context.Background()
	ch := make(chan MultiReturns)
	go func() {
		r, err := Waterfall(rootCtx, func(ctx context.Context, n interface{}) (interface{}, error) {
			time.Sleep(100 * time.Millisecond)
			return 1, nil
		}, func(ctx context.Context, n interface{}) (interface{}, error) {
			time.Sleep(100 * time.Millisecond)
			n = n.(int) + 1
			return n, nil
		}, func(ctx context.Context, n interface{}) (interface{}, error) {
			time.Sleep(100 * time.Millisecond)
			n = n.(int) + 1
			return n, nil
		})

		ch <- MultiReturns{r, err}
		close(ch)
	}()

	r := <-ch
	assert.Nil(t, r.Err)

	assert.IsTypef(t, 1, r.Val, "Should return int, but got: %v", r.Val)
	assert.Equal(t, 3, r.Val)
}

func TestRetry_ShouldReturnsOnSuccess(t *testing.T) {
	t.Parallel()

	rootCtx := context.Background()
	ch := make(chan MultiReturns)
	go func() {
		r, err := Retry(rootCtx, 3, func(ctx context.Context) (interface{}, error) {
			time.Sleep(200 * time.Millisecond)
			return 1, nil
		})

		ch <- MultiReturns{r, err}
		close(ch)
	}()

	r := <-ch
	assert.Nil(t, r.Err)
	assert.Equal(t, 1, r.Val)
}

func TestRetry_ShouldReturnsAfterFailingSeveralTimes(t *testing.T) {
	t.Parallel()

	var times int32
	rootCtx := context.Background()
	ch := make(chan MultiReturns)
	go func() {
		r, err := Retry(rootCtx, 3, func(ctx context.Context) (interface{}, error) {
			atomic.AddInt32(&times, 1)
			if atomic.LoadInt32(&times) >= 2 {
				return 1, nil
			}
			return nil, fmt.Errorf("err")
		})

		ch <- MultiReturns{r, err}
		close(ch)
	}()

	r := <-ch
	assert.Nil(t, r.Err)

	assert.Equal(t, int32(2), times)

	assert.Equal(t, 1, r.Val)
}

func TestRetry_ShouldKeepRetrying(t *testing.T) {
	t.Parallel()

	var times int32
	rootCtx := context.Background()
	ch := make(chan MultiReturns)
	go func() {

		r, err := Retry(rootCtx, 3, func(ctx context.Context) (interface{}, error) {
			atomic.AddInt32(&times, 1)
			return nil, fmt.Errorf("err")
		})

		ch <- MultiReturns{r, err}
		close(ch)
	}()

	r := <-ch

	assert.NotNil(t, r.Err)

	assert.Equal(t, int32(3), times)
}

func TestRetry_WhenRootCtxCanceled(t *testing.T) {
	t.Parallel()

	rootCtx, cancel := context.WithCancel(context.Background())
	ch := make(chan MultiReturns)
	go func() {

		r, err := Retry(rootCtx, 3, func(ctx context.Context) (interface{}, error) {
			time.Sleep(50 * time.Millisecond)
			return nil, fmt.Errorf("err")
		})

		ch <- MultiReturns{r, err}
		close(ch)
	}()
	go func() {
		time.Sleep(10 * time.Millisecond)
		cancel()
	}()

	r := <-ch

	assert.NotNil(t, r.Err)
}

func TestThrow_ShouldReturnParentContextErrorWhenParentContextIsDone(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	cancel() // Cancel the context immediately

	err := Throw(ctx, func(ctx context.Context) (int, error) {
		time.Sleep(100 * time.Millisecond)
		return 1, nil
	})

	assert.IsTypef(t, context.Canceled, err, "Should return context.Canceled, but got: %v", err)
}

func TestThrow_ShouldReturnErrorWhenExecutableFails(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	err := Throw(ctx, func(ctx context.Context) (int, error) {
		return 0, fmt.Errorf("executable error")
	})
	assert.NotNil(t, err)
}

func TestThrow_ShouldReturnNilWhenAllExecutablesSucceed(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	err := Throw(ctx, func(ctx context.Context) (int, error) {
		return 1, nil
	}, func(ctx context.Context) (int, error) {
		return 2, nil
	}, func(ctx context.Context) (int, error) {
		return 3, nil
	})

	if err != nil {
		t.Errorf("Expected nil, but got: %v", err)
	}
}

func TestMin(t *testing.T) {
	t.Parallel()

	if min(1, 2) != 1 {
		t.Errorf("Should returns 1")
	}
	if min(-1, 2) != -1 {
		t.Errorf("Should returns -1")
	}
	if min(54321, 12345) != 12345 {
		t.Errorf("Should returns 12345")
	}
}
