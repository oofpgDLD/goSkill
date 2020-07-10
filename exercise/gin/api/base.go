package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

func New() *gin.Engine{
	//config
	f,err := os.OpenFile("gin.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	//f, err := os.Create("gin.log")
	if err != nil {
		panic(err)
	}
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	//init gin
	r := gin.Default()

	root := r.Group("/gin")
	root.GET("/user", User)
	root.Any("/map", Map)
	root.GET("/JSONP?callback=x", JSONP)
	return r
}

func User(c *gin.Context) {
	u := make(map[string]interface{})
	u["name"] = "test"
	u["id"] = 1

	/*b, err := json.Marshal(u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}*/
	c.JSON(http.StatusOK, u)
}

func Map(c *gin.Context) {
	//url params
	ids := c.QueryMap("ids")
	array := c.QueryArray("array")
	//topic := c.Query("topic")
	topic := c.DefaultQuery("topic", "default topic")
	c.MustBindWith()
	//post body params
	names := c.PostFormMap("names")

	ret := struct {
		Ids interface{} `json:"ids"`
		Names interface{} `json:"names"`
		Array interface{} `json:"array"`
		Topic interface{} `json:"topic"`
	}{
		Ids: ids,
		Names: names,
		Array: array,
		Topic: topic,
	}
	fmt.Printf("ids: %v; names: %v", ids, names)
	c.JSON(http.StatusOK, ret)
}

func JSONP(c *gin.Context) {
	data := map[string]interface{}{
		"foo": "bar",
	}

	//callback is x
	// Will output  :   x({\"foo\":\"bar\"})
	c.JSONP(http.StatusOK, data)
}