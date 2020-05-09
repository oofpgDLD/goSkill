package channel

import (
	"fmt"
	"github.com/limetext/backend/log"
	"testing"
	"time"
)

//通道
func Test_Channel(t *testing.T) {
	ch := make(chan int, 200)
	ch2 := make(chan string)
	//to string

	go func() {
		for x := range ch{
			time.Sleep(time.Second)
			ch2 <- fmt.Sprintf("%v", x)
		}
		close(ch2)
		log.Info("传递结束")
	}()

	go func() {
		for i:= 1 ; i <= 60; i++{
			ch <- i
		}
		close(ch)
		log.Info("输入结束")
	}()

	for x := range ch2{
		fmt.Println(x)
	}
	log.Info("打印结束")
}