package main

import (
	"encoding/json"
	"fmt"
	"database/sql"
	"log"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)


type GetImage struct{
	ImageID		string `json:"imageID"`
}

type AnnotateImage struct{
	ImageID		string `json:"imageID"`
	Q1			string `json:"Q1"`
	Q2			string `json:"Q2"`
	Q3			string `json:"Q3"`
	Q4			string `json:"Q4"`
	Q5			string `json:"Q5"`
	Q6			string `json:"Q6"`
}

func getDBConnect() (*sql.DB, error){
	db, err := sql.Open("mysql", "root:Jxyz30932!@(127.0.0.1:3306)/comp2002")

	if err != nil {
		log.Println("Connect to DB error:", err.Error())
	}else{
		db.SetConnMaxLifetime(7 * time.Second)
		db.SetMaxOpenConns(10)
		db.SetConnMaxIdleTime(10)
	}
	return db, err
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

func saveRecord(annot AnnotateImage) bool{
	var flag bool = false
	db, err := getDBConnect() // 连接到数据库

	// r, err := db.Query("SELECT * FROM `init_test` WHERE sImageID = ?", annot.ImageID) 
	// if (err == nil){
		_, err = db.Exec("INSERT INTO init_test (sQ1, sQ2, sQ3, sQ4, sQ5, sQ6, sImageID) VALUES (?, ?, ?, ?, ?, ?, ?) ON DUPLICATE KEY UPDATE" + 
		" sQ1 = ?, sQ2 = ?, sQ3 = ?, sQ4 = ?, sQ5 = ?, sQ6 = ?, sImageID = ?;", 
						annot.Q1, annot.Q2, annot.Q3, 
						annot.Q4, annot.Q5, annot.Q6,
						annot.ImageID, annot.Q1, annot.Q2, annot.Q3, 
						annot.Q4, annot.Q5, annot.Q6,
						annot.ImageID) // 操作成功返回true, 否则打印错误消息并返回false   sQ1 = ?, sQ2 = ?, sQ3 = ?, sQ4 = ?, sQ5 = ?, sQ6 = ?, sImageID = ?
		if err != nil{
			log.Println("Do record saving error:", err.Error())
			flag = false
		}

	flag = true
	defer db.Close()
	
	// 确保连接可用
	err = db.Ping()
	if err != nil{
		log.Println("Do db.Ping() error:", err.Error())
	}

	return flag
}

func getRecord(id string) string{
	
	db, err := getDBConnect()

	if err != nil{
		log.Println("Do get record fail:", err.Error())
	}

	rows, err := db.Query("SELECT * FROM `init_test` WHERE sImageID = ?", id)
	if err != nil{
		log.Println("Do query error: ", err)
		return ""
	}

	defer rows.Close()
	var result string
	for rows.Next(){
		var image AnnotateImage
		err = rows.Scan(&image.Q1, &image.Q2, &image.Q3, 
			&image.Q4, &image.Q5, &image.Q6, &image.ImageID)
		
		if err != nil{
			panic(err.Error())
		}

		mashal, err := json.Marshal(&image)
		if err != nil{
			log.Println("An error has occurred: ", err)
			return ""
		}
		result = string(mashal)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil{
		log.Println("Do db.Ping() error: ", err.Error())
	}

	return result
}

func main(){
	r := gin.Default()
	r.Use(Cors())

	// Save the record to DB
	r.POST("/save", func(c *gin.Context){
		var imgAnnot AnnotateImage
		if err := c.ShouldBind(&imgAnnot);err != nil{
			fmt.Println("Fail")
		}

		resultFlag := saveRecord(imgAnnot)
		if true == resultFlag{
			c.JSON(http.StatusOK, gin.H{
				"message": "Record saved successfully!",
			})
		}else{
			c.JSON(http.StatusExpectationFailed, gin.H{
				"message": "Record saved failed!",
			})
		}
	})

	// Display the data fetched from DB
	r.GET("/get", func(c *gin.Context) {
		id := c.Query("imageID")
		
		resultInfo := getRecord(id)
		var data AnnotateImage
		if err := json.Unmarshal([]byte(resultInfo), &data); err == nil{
			fmt.Println(data)
		}else{
			fmt.Println(err)
		}
		c.JSON(http.StatusOK, data)
	})

	r.Run(":8080")
}
