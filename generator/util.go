package generator

import "context"

func ToChannelSlice[V any, Vs []V](ctx context.Context, val Vs) (i chan int, v chan V) {
	i = make(chan int, len(val))
	v = make(chan V, len(val))
	go ChannelSlice(ctx, i, v, val)
	return
}

func ToChannelMap[I comparable, V any, Vs ~map[I]V](ctx context.Context, val Vs) (i chan I, v chan V) {
	i = make(chan I, len(val))
	v = make(chan V, len(val))
	go ChannelMap(ctx, i, v, val)

	return
}

func ChannelSlice[V any, Vs ~[]V](c context.Context, i chan int, r chan V, val Vs) {
OUTERLOOP:
	for {
		select {
		case <-c.Done():
			break OUTERLOOP
		default:
			for k, l := range val {
				i <- k
				r <- l
			}
			break OUTERLOOP
		}
	}

	close(i)
	close(r)
}

func ChannelMap[I comparable, V any, Vs ~map[I]V](c context.Context, i chan I, r chan V, val Vs) {
OUTERLOOP:
	for {
		select {
		case <-c.Done():
			break OUTERLOOP
		default:
			for k, l := range val {
				i <- k
				r <- l
			}
			break OUTERLOOP
		}
	}

	close(i)
	close(r)
}
