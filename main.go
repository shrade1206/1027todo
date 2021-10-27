package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Todo struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

var upgrader = &websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
var todo Todo

func main() {
	initMysql()
	defer DB.Close()

	r := gin.Default()
	r.POST("/insert", func(c *gin.Context) {
		c.BindJSON(&todo)
		// err := c.BindJSON(&todo)
		// if err != nil {
		// 	log.Printf("BindJson Error : %s", err.Error())
		// }
		err := DB.Create(&todo).Error
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"Title": &todo,
			})
		}
	})
	r.GET("/get", func(c *gin.Context) {
		var todos []Todo
		if err := DB.Find(&todos).Error; err != nil {
			c.JSON(http.StatusOK, gin.H{
				"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, todos)
		}
	})

	r.NoRoute(gin.WrapH(http.FileServer(http.Dir("./public"))), func(c *gin.Context) {
		path := c.Request.URL.Path
		method := c.Request.Method
		fmt.Println(path)
		fmt.Println(method)
		//檢查path的開頭使是否為"/"
		if strings.HasPrefix(path, "/") {
			fmt.Println("ok")
		}
	})
	err := r.Run(":8080")
	if err != nil {
		log.Fatal("8080 err : ", err.Error())
	}
}

// r.POST("/insert", func(c *gin.Context) {
// 	c.BindJSON(&todo)
// 	// err := c.BindJSON(&todo)
// 	// if err != nil {
// 	// 	log.Printf("BindJson Error : %s", err.Error())
// 	// }
// 	err := DB.Create(&todo).Error
// 	if err != nil {
// 		c.JSON(http.StatusOK, gin.H{
// 			"error": err.Error()})
// 	} else {
// 		c.JSON(http.StatusOK, gin.H{
// 			"Title": todo.Title,
// 		})
// 		// c.JSON(http.StatusOK, gin.H{
// 		// 	"code": 2000,
// 		// 	"msg":  "success",
// 		// 	"data": todo,
// 		// })
// 	}
// })
