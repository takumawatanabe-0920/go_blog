package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type JsonRequest struct {
	FieldStr  string `json:"field_str"`
	FieldInt  int    `json:"field_int"`
	FieldBool bool   `json:"field_bool"`
}

func main() {
	  router := gin.Default()
    /* 
		  要件1: find getById post put delete
			要件2: queryを受け取れるようにする
			要件3: レスポンスをjsonで返す
		*/

		router.GET("/blog", func(c *gin.Context) {
      blogType := c.Query("type")
			fmt.Println(blogType)
			c.JSON(200, gin.H{
				"message": "blog",
			})
		})

		router.GET("/blog/:id", func(c *gin.Context) {
			id := c.Param("id")
			fmt.Printf("%v\n", id)
			c.JSON(200, gin.H{
				"message": "blogById",
			})
		})

		router.POST("/blog", func(c *gin.Context) {
			var json JsonRequest
			if err := c.ShouldBindJSON(&json); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"str": json.FieldStr, "int": json.FieldInt, "bool": json.FieldBool})
		})

		router.PUT("/blog/:id", func(c *gin.Context) {
			id := c.Param("id")
			fmt.Println(id)
			var json JsonRequest
			if err := c.ShouldBindJSON(&json); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.fStatusOK, gin.H{"str": json.FieldStr, "int": json.FieldInt, "bool": json.FieldBool})
		})

		router.DELETE("/blog/:id", func(c *gin.Context) {
			id := c.Param("id")
			fmt.Println(id)
			c.JSON(200, gin.H{
				"message": "blogDelete",
			})
		})

	  router.Run()
}