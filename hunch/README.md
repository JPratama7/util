
## Fork from github.com/aaronjan/hunch


## Usage

### Installation

#### `go get`

```shell
$ go get -u -v github.com/aaronjan/hunch
```

#### `go mod` (Recommended)

```go
import "github.com/aaronjan/hunch"
```

```shell
$ go mod tidy
```

### Types

```go
type Executable func(context.Context) (interface{}, error)

type ExecutableInSequence func(context.Context, interface{}) (interface{}, error)
```

### API

#### All

```go
func All(parentCtx context.Context, execs ...Executable) ([]interface{}, error) 
```

All returns all the outputs from all Executables, order guaranteed.

##### Examples

```go
ctx := context.Background()
r, err := hunch.All(
    ctx,
    func(ctx context.Context) (interface{}, error) {
        time.Sleep(300 * time.Millisecond)
        return 1, nil
    },
    func(ctx context.Context) (interface{}, error) {
        time.Sleep(200 * time.Millisecond)
        return 2, nil
    },
    func(ctx context.Context) (interface{}, error) {
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
func Take(parentCtx context.Context, num int, execs ...Executable) ([]interface{}, error)
```

Take returns the first `num` values outputted by the Executables.

##### Examples

```go
ctx := context.Background()
r, err := hunch.Take(
    ctx,
    // Only need the first 2 values.
    2,
    func(ctx context.Context) (interface{}, error) {
        time.Sleep(300 * time.Millisecond)
        return 1, nil
    },
    func(ctx context.Context) (interface{}, error) {
        time.Sleep(200 * time.Millisecond)
        return 2, nil
    },
    func(ctx context.Context) (interface{}, error) {
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
func Last(parentCtx context.Context, num int, execs ...Executable) ([]interface{}, error)
```

Last returns the last `num` values outputted by the Executables.

##### Examples

```go
ctx := context.Background()
r, err := hunch.Last(
    ctx,
    // Only need the last 2 values.
    2,
    func(ctx context.Context) (interface{}, error) {
        time.Sleep(300 * time.Millisecond)
        return 1, nil
    },
    func(ctx context.Context) (interface{}, error) {
        time.Sleep(200 * time.Millisecond)
        return 2, nil
    },
    func(ctx context.Context) (interface{}, error) {
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
func Waterfall(parentCtx context.Context, execs ...ExecutableInSequence) (interface{}, error)
```

Waterfall runs `ExecutableInSequence`s one by one, passing previous result to next Executable as input. When an error occurred, it stop the process then returns the error. When the parent Context canceled, it returns the `Err()` of it immediately.

##### Examples

```go
ctx := context.Background()
r, err := hunch.Waterfall(
    ctx,
    func(ctx context.Context, n interface{}) (interface{}, error) {
        return 1, nil
    },
    func(ctx context.Context, n interface{}) (interface{}, error) {
        return n.(int) + 1, nil
    },
    func(ctx context.Context, n interface{}) (interface{}, error) {
        return n.(int) + 1, nil
    },
)

fmt.Println(r, err)
// Output:
// 3 <nil>
```

#### Retry

```go
func Retry(parentCtx context.Context, retries int, fn Executable) (interface{}, error)
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
r, err := hunch.Retry(
    ctx,
    10,
    func(ctx context.Context) (interface{}, error) {
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
