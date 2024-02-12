package hunch

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type mockExecutable struct {
	value int
	delay time.Duration
}

func (m mockExecutable) Execute(ctx context.Context) (int, error) {
	time.Sleep(m.delay)
	return m.value, nil
}

type mockExecutableFail struct {
	err error
}

func (m mockExecutableFail) Execute(ctx context.Context) (int, error) {
	return 0, errors.New("error")
}

func TestFirstReturnsFirstResolvedValue(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	exec1 := mockExecutable{value: 1, delay: 2 * time.Second}
	exec2 := mockExecutable{value: 2, delay: 1 * time.Second}
	val, err := First(ctx, exec1.Execute, exec2.Execute)

	assert.Nil(t, err)
	assert.Equal(t, exec2.value, val)
}

func TestFirstReturnsErrorWhenAllExecutablesFail(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	exec1 := mockExecutable{value: 1, delay: 2 * time.Second}
	exec2 := mockExecutable{value: 2, delay: 1 * time.Second}
	val, err := First(ctx, exec1.Execute, exec2.Execute)

	assert.Nil(t, err)
	assert.Equal(t, val, exec2.value)
}

func TestFirstCancelsWhenContextDone(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	exec1 := mockExecutable{value: 1, delay: 2 * time.Second}
	exec2 := mockExecutable{value: 2, delay: 4 * time.Second}
	_, err := First(ctx, exec1.Execute, exec2.Execute)
	assert.Nil(t, err)
}

func TestThrowExecutesAllExecutablesIgnoringOutputs(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	exec1 := mockExecutable{value: 1, delay: 1 * time.Second}
	exec2 := mockExecutable{value: 2, delay: 1 * time.Second}
	Throw(ctx, exec1.Execute, exec2.Execute)
}

func TestThrowReturnsErrorWhenAnyExecutableFails(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	exec1 := mockExecutable{value: 1, delay: 1 * time.Second}
	exec2 := mockExecutable{value: 2, delay: 1 * time.Second}
	Throw(ctx, exec1.Execute, exec2.Execute)
}

func TestThrowCancelsWhenContextDone(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	exec1 := mockExecutable{value: 1, delay: 2 * time.Second}
	exec2 := mockExecutable{value: 2, delay: 3 * time.Second}
	Throw(ctx, exec1.Execute, exec2.Execute)
}

func TestAllReturnsAllOutputsInOrder(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	exec1 := mockExecutable{value: 1, delay: 1 * time.Second}
	exec2 := mockExecutable{value: 2, delay: 2 * time.Second}
	exec3 := mockExecutable{value: 3, delay: 3 * time.Second}
	values, err := All(ctx, exec1.Execute, exec2.Execute, exec3.Execute)

	assert.Nil(t, err)
	assert.Equal(t, []int{1, 2, 3}, values)
}

func TestAllReturnsErrorWhenAnyExecutableFails(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	exec1 := mockExecutable{value: 1, delay: 1 * time.Second}
	exec2 := mockExecutableFail{}
	values, err := All(ctx, exec1.Execute, exec2.Execute)

	assert.NotNil(t, err)
	assert.Nil(t, values)
}

func TestAllCancelsWhenContextDone(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	exec1 := mockExecutable{value: 1, delay: 1 * time.Second}
	exec2 := mockExecutable{value: 2, delay: 3 * time.Second}
	values, err := All(ctx, exec1.Execute, exec2.Execute)

	assert.NotNil(t, err)
	assert.Nil(t, values)
}

func TestTakeReturnsRequestedNumberOfValues(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	exec1 := mockExecutable{value: 1, delay: 1 * time.Second}
	exec2 := mockExecutable{value: 2, delay: 1 * time.Second}
	exec3 := mockExecutable{value: 3, delay: 1 * time.Second}
	values, err := Take(ctx, 2, exec1.Execute, exec2.Execute, exec3.Execute)

	assert.Nil(t, err)
	assert.Equal(t, []int{1, 2}, values)
}

func TestTakeReturnsAllValuesIfRequestedMoreThanAvailable(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	exec1 := mockExecutable{value: 1, delay: 1 * time.Second}
	exec2 := mockExecutable{value: 2, delay: 1 * time.Second}
	values, err := Take(ctx, 3, exec1.Execute, exec2.Execute)

	assert.Nil(t, err)
	assert.Equal(t, []int{1, 2}, values)
}

func TestTakeReturnsErrorWhenAnyExecutableFails(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	exec1 := mockExecutable{value: 1, delay: 1 * time.Second}
	exec2 := mockExecutableFail{}
	values, err := Take(ctx, 2, exec1.Execute, exec2.Execute)

	assert.NotNil(t, err)
	assert.Nil(t, values)
}

func TestTakeCancelsWhenContextDone(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	exec1 := mockExecutable{value: 1, delay: 1 * time.Second}
	exec2 := mockExecutable{value: 2, delay: 3 * time.Second}
	values, err := Take(ctx, 2, exec1.Execute, exec2.Execute)

	assert.NotNil(t, err)
	assert.Nil(t, values)
}
