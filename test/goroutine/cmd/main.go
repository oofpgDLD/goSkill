package main

import (
	"fmt"
	"time"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	signal := make(chan bool)

	wg.Add(2)
	//func a
	go func() {
		for {
			select {
			case rlt, ok := <-signal:
				if rlt && ok {
					for i :=0 ;i<= 100; i++ {
						fmt.Println("a")
					}
					wg.Done()
					return
				}
			}
		}
	}()

	//func b
	go func() {
		for {
			select {
			case rlt, ok := <-signal:
				if rlt && ok {
					for i :=0 ;i<= 100; i++ {
						fmt.Println("b")
					}
					wg.Done()
					return
				}
			}
		}
	}()

	time.Sleep(time.Second * 2)
	signal <- true
	signal <- true

	wg.Wait()
	fmt.Println("exit...")
}

