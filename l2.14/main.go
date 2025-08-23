package main

import (
	"fmt"
	"time"
)

func or(channels ...<-chan interface{}) <-chan interface{} {
	if len(channels) == 0 {
		return nil
	}
	if len(channels) == 1 {
		return channels[0]
	}

	out := make(chan interface{})
	go func() {
		defer close(out)
		select {
		case <-channels[0]:
			return
		case <-or(channels[1:]...):
			return
		}
	}()
	return out
}

// func or(channels ...<-chan interface{}) <-chan interface{} {
// 	out := make(chan interface{})
// 	go func() {
// 		defer close(out)
// 		for {
// 			select {
// 			case <-out:
// 				return
// 			default:
// 				for _, ch := range channels {
// 					select {
// 					case <-ch:
// 						return
// 					default:
// 					}
// 				}
// 			}
// 		}
// 	}()
// 	return out
// }

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(3*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	fmt.Println("done after:", time.Since(start))
}
