// package main

// import (
// 	"github.com/gin-gonic/gin"
// 	// "log"
// 	"fmt"
// 	// "database/sql"
// 	"net/http"
// )

// func main(){
// 	r := gin.Default()

// 	r.GET("/", func(c *gin.Context){
// 		params := c.Request.URL.Query()
// 		fmt.Println("Params:", params)

// 		c.JSON(http.StatusOK, gin.H{
// 			"message": "Fetched",
// 		})
// 	})

// 	r.Run(":8888")
// }