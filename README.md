# ASKAA
 为创新实践写的后端（只有增删改查）

## 后端接口说明

程序说明，就是一个很简单的对于

本程序使用**RESTful**编写

从代码可以看出

```go
func main() {
	r := gin.Default()
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
```

统一使用`/question`作为应用程序接口

在GET方法中，使用Query String

使用方法：http://localhost:9090/question?question=  +(问题)

在POST方法中，允许前端传入JSON，我将JSON解析，进行数据的插入

在PUT方法中，通过ID的查询，通过JSON的解析，进行数据的更新

在DELETE方法中：使用 http://localhost:9090/question/(ID号)

进行数据的删除（软删除）因为存在`gorm.Model`字段，所以删除只是在`DeleteAt`添加一个时间，同时保证不会被代码读取

## 使用方法

首先需要在MySQL数据库中建立一个数据库，然后通过mysql.go这个文件，将`Username`,`Password`,`hostname`,`dbname`填写好以后，在安装好`golang`的机器上面，运行如下命令

```shell
go mod tidy
cd src
go build
然后运行生成的可执行文件，一般是文件夹的名字
用Nginx反向代理运行就可以了
（默认端口号:9090）
```



**感谢E99p1ant小哥哥的倾力支持**

