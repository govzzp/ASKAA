package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	r := gin.Default()
	r.Use(Cors())
	question := r.Group("/question")
	{
		question.GET("/",GetQuestion)
		question.POST("/",CreateQuestion)
		question.PUT("/:id",UpdateQuestion)
		question.DELETE("/:id",DeleteQuestion)
	}
	err := r.Run(":9090")
	if err != nil {
		fmt.Printf("Start Service Error err:%#v\n",err)
		return
	}
}
func GetQuestion(c *gin.Context){
	question := new(AnswerSheet)
	questions := c.Query("question")
	db.Model(&AnswerSheet{}).Where(&AnswerSheet{Question: questions}).Find(question)
	c.JSON(http.StatusOK, question)
}
func CreateQuestion(c *gin.Context ){


	var input struct {
		Question string `json:"question";binding:"required"`
		Answer   string `json:"answer";binding:"required"`
		Depart   string `json:"depart";binding:"required"`
		Name     string `json:"name";binding:"required"`
	}
	err := c.BindJSON(&input)
	if err != nil{
		c.JSON(400, gin.H{
			"message": "error payload",
		})
		return
	}

	question := AnswerSheet{
		Question: input.Question,
		Answer: input.Answer,
		Depart: input.Depart,
		Name: input.Name,
	}

	tx := db.Begin()
	if tx.Create(&question).RowsAffected != 1{
		db.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"message" : "database error",
		})
		return
	}
	tx.Commit()

	c.JSON(http.StatusCreated,gin.H{
		"status":http.StatusCreated,
		"message":"Question insert ok",
		"message_id":question.ID,
	})
}
func UpdateQuestion(c *gin.Context){
	var input struct {
		Question string `json:"question";binding:"required"`
		Answer   string `json:"answer";binding:"required"`
		Depart   string `json:"depart";binding:"required"`
		Name     string `json:"name";binding:"required"`
		ID 		 int`json:"id";binding:"required"`
	}
	err := c.BindJSON(&input)
	if err != nil{
		c.JSON(400, gin.H{
			"message": "error update",
		})
		return
	}
	question := AnswerSheet{
		Question: input.Question,
		Answer: input.Answer,
		Depart: input.Depart,
		Name: input.Name,
	}
	tx := db.Begin()
	if tx.Update(&question).RowsAffected != 1{
		db.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"message" : "database error",
		})
		return
	}
	tx.Commit()

	c.JSON(http.StatusUpgradeRequired,gin.H{
		"status":http.StatusUpgradeRequired,
		"message":"Question update ok",
		"message_id":question.ID,
	})
}
func DeleteQuestion(c *gin.Context)  {
	cid := c.Param("id")
	id , err := strconv.Atoi(cid)
	if err != nil {
		c.JSON(http.StatusNotFound,gin.H{
			"message":"Delete Error",
		})
		return
	}
	db.Delete(&AnswerSheet{},"id = ?",id)
	c.JSON(http.StatusOK,gin.H{
		"message":"Delete Successful",
	})
}
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		var headerKeys []string
		for k , _ :=range c.Request.Header {
			headerKeys = append(headerKeys,k)
		}
		headerStr := strings.Join(headerKeys,",")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*")        // 这是允许访问所有域
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")      //服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
			//  header的类型
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			//              允许跨域设置                                                                                                      可以返回其他子段
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar")      // 跨域关键设置 让浏览器可以解析
			c.Header("Access-Control-Max-Age", "172800")        // 缓存请求信息 单位为秒
			c.Header("Access-Control-Allow-Credentials", "false")       //  跨域请求是否需要带cookie信息 默认设置为true
			c.Set("content-type", "application/json")       // 设置返回格式是json
		}

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		// 处理请求
		c.Next()        //  处理请求
	}
}

