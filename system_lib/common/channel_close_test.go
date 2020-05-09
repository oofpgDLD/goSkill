package common

import (
	"fmt"
	"testing"
	"time"
)

func Test_CloseChannel(t *testing.T) {
	closer := make(chan interface{})
	ch := make(chan int)

	go func() {
		for i:=1;;i++{
			select {
			case <-closer:
				return
			default:
				time.Sleep(time.Millisecond * 500)
				ch <- i

				if i == 20 {
					close(closer)
					close(ch)
				}
			}
		}
	}()

	for {
		select {
		case i := <-ch:
			time.Sleep(time.Second * 1)
			fmt.Println(i)
		default:
		}
	}
}