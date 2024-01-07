
## Fork from github.com/aaronjan/hunch


## Usage

### Installation

#### `go get`

```shell
$ go get -u -v github.com/JPratama7/util/hunch
```

#### `go mod` (Recommended)

```go
import "github.com/JPratama7/util/hunch"
```

```shell
$ go mod tidy
```

### Types

```go
type Executable[T any] func(context.Context) (T, error)

type ExecutableInSequence[T any] func(context.Context, T) (T, error)
```

### API

#### All

```go
func All[T any](parentCtx context.Context, execs ...Executable[T]) ([]T, error)
```

All returns all the outputs from all Executables, order guaranteed.

##### Examples

```go
ctx := context.Background()
r, err := hunch.All[int](
    ctx,
    func(ctx context.Context) (int, error) {
        time.Sleep(300 * time.Millisecond)
        return 1, nil
    },
    func(ctx context.Context) (int, error) {
        time.Sleep(200 * time.Millisecond)
        return 2, nil
    },
    func(ctx context.Context) (int, error) {
        time.Sleep(100 * time.Millisecond)
        return 3, nil
    },
)

fmt.Println(r, err)
// Output:
// [1 2 3] <nil>
```

#### Take

```go
func Take[T any](parentCtx context.Context, num int, execs ...Executable[T]) ([]T, error)
```

Take returns the first `num` values outputted by the Executables.

##### Examples

```go
ctx := context.Background()
r, err := hunch.Take[int](
    ctx,
    // Only need the first 2 values.
    2,
    func(ctx context.Context) (int, error) {
        time.Sleep(300 * time.Millisecond)
        return 1, nil
    },
    func(ctx context.Context) (int, error) {
        time.Sleep(200 * time.Millisecond)
        return 2, nil
    },
    func(ctx context.Context) (int, error) {
        time.Sleep(100 * time.Millisecond)
        return 3, nil
    },
)

fmt.Println(r, err)
// Output:
// [3 2] <nil>
```

#### Last

```go
func Last[T any](parentCtx context.Context, num int, execs ...Executable[T]) ([]T, error)
```

Last returns the last `num` values outputted by the Executables.

##### Examples

```go
ctx := context.Background()
r, err := hunch.Last[int](
    ctx,
    // Only need the last 2 values.
    2,
    func(ctx context.Context) (int, error) {
        time.Sleep(300 * time.Millisecond)
        return 1, nil
    },
    func(ctx context.Context) (int, error) {
        time.Sleep(200 * time.Millisecond)
        return 2, nil
    },
    func(ctx context.Context) (int, error) {
        time.Sleep(100 * time.Millisecond)
        return 3, nil
    },
)

fmt.Println(r, err)
// Output:
// [2 1] <nil>
```

#### Waterfall

```go
func Waterfall[T any](parentCtx context.Context, execs ...ExecutableInSequence[T]) (T, error)
```

Waterfall runs `ExecutableInSequence`s one by one, passing previous result to next Executable as input. When an error occurred, it stop the process then returns the error. When the parent Context canceled, it returns the `Err()` of it immediately.

##### Examples

```go
ctx := context.Background()
r, err := hunch.Waterfall[int](
    ctx,
    func(ctx context.Context, n int) (int, error) {
        return 1, nil
    },
    func(ctx context.Context, n int) (int, error) {
        return n.(int) + 1, nil
    },
    func(ctx context.Context, n int) (int, error) {
        return n.(int) + 1, nil
    },
)

fmt.Println(r, err)
// Output:
// 3 <nil>
```

#### Retry

```go
func Retry[T any](parentCtx context.Context, retries int, fn Executable[T]) (T, error)
```

Retry attempts to get a value from an Executable instead of an Error. It will keeps re-running the Executable when failed no more than `retries` times. Also, when the parent Context canceled, it returns the `Err()` of it immediately.

##### Examples

```go
count := 0
getStuffFromAPI := func() (int, error) {
    if count == 5 {
        return 1, nil
    }
    count++

    return 0, fmt.Errorf("timeout")
}

ctx := context.Background()
r, err := hunch.Retry[int](
    ctx,
    10,
    func(ctx context.Context) (int, error) {
        rs, err := getStuffFromAPI()

        return rs, err
    },
)

fmt.Println(r, err, count)
// Output:
// 1 <nil> 5
```

## Credits

Heavily inspired by [Async](https://github.com/caolan/async/) and [ReactiveX](http://reactivex.io/).

## Licence

[Apache 2.0](https://www.apache.org/licenses/LICENSE-2.0)
