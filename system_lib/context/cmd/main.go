package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func MidWare(c *gin.Context) (){
	ctx, _ := context.WithTimeout(c.Request.Context(), time.Second * 3)
	//defer cancel()
	c.Request = c.Request.WithContext(ctx)
}

func MidWare2(c *gin.Context) (){
	go func() {
		time.Sleep(time.Second * 1)
		c.JSON(http.StatusOK,"{\"key\":1}")
	}()
}

func Test2(c *gin.Context) (){
	ch := make(chan bool)
	go func() {
		time.Sleep(time.Second * 1)
		c.Set("result",  3)
		ch<- true
	}()

	go func() {
		select {
		case <-c.Request.Context().Done():
			ch <- false
			return
		}
	}()
	/*i := 0
	for {
		select {
		case <-c.Request.Context().Done():
			fmt.Println("done")
			return
		default:
		}
		i++
		time.Sleep(time.Second * 1)
		fmt.Println("do:", i)
	}*/

	select {
	case b := <-ch:
		fmt.Println("done")
		if b {
			c.String(http.StatusOK,"", c.MustGet("result").(int))
			return
		}else {
			fmt.Println("time out")
		}
		c.String(http.StatusOK,"{}")
		return
	}
}

func Test(c *gin.Context) (){
	/*i := 0
	for {
		select {
		case <-c.Request.Context().Done():
			fmt.Println("done")
			return
		default:
		}
		i++
		time.Sleep(time.Second * 1)
		fmt.Println("do:", i)
	}*/
	time.Sleep(time.Second * 100)
	c.String(http.StatusOK,"", "{}")
}

func main() {
	root := gin.Default()
	root.GET("/hello", MidWare, Test)
	root.GET("/hello2", MidWare, Test2)

	root.Run("localhost:8999")
}
