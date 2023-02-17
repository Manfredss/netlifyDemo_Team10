package main

import (
	"github.com/gin-gonic/gin"
	// "log"
	"fmt"
	// "encoding/json"
	"net/http"
)

type AnnotateImage struct{
	Q1			string `json:"Q1"`
	Q2			string `json:"Q2"`
	Q3			string `json:"Q3"`
	Q4			string `json:"Q4"`
	Q5			string `json:"Q5"`
	Q6			string `json:"Q6"`
	imageid		string `json:"imageid"`
}

func main(){
	r := gin.Default()
	r.Use(Cors())
	// http.HandleFunc("/hi", handlePostJson)
	r.POST("/hi", func(c *gin.Context){
		var imgAnnot AnnotateImage
		if err := c.ShouldBind(&imgAnnot);err != nil{
			fmt.Println("Fail")
		}
		// Q1 := c.PostForm("Q1")
		// _ = c.Request.ParseForm()
		// fmt.Println(c.Request.PostForm)
		// fmt.Println(imgAnnot)
		fmt.Println("Q1", imgAnnot.Q1)
		fmt.Println("Q2", imgAnnot.Q2)
		fmt.Println("Q3", imgAnnot.Q3)
		fmt.Println("Q4", imgAnnot.Q4)
		fmt.Println("Q5", imgAnnot.Q5)
		fmt.Println("Q6", imgAnnot.Q6)
		fmt.Println("imageid", imgAnnot.imageid)
		// fmt.Println(imgAnnot.Q2)
		// c.JSON(http.StatusOK, c.Request.PostForm)
			// gin.H{"message:": "Received",}
		// )
	})

	r.Run(":8080")
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		c.Header("Access-Control-Allow-Origin", origin)
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token,X-Token,X-User-Id")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS,DELETE,PUT")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}

// func handlePostJson(writer http.ResponseWriter, request *http.Request) {
// 	decoder := json.NewDecoder(request.Body)

// 	var params map[string]string

// 	decoder.Decode(params)

// 	fmt.Println("%s, %s, %s", params["imageID"], params["Q1"], params["Q2"])


// }