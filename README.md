# ASKAA
 为创新实践写的后端（只有增删改查）

## 相关技术说明

* **Golang：一门新的程序设计语言**
* **Gin：Golang的一种框架**
* **GORM：对于Golang语言友好的一种开发人员ORM库**
* **ORM：一种对象关系映射，用来将对象和数据库之间的映射的元数据，将面向对象语言程序中的对象自动持久化到关系数据库中。 本质上就是将数据从一种形式转换到另外一种形式。**

## 程序说明

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

##  后端接口说明

### 系统接口：

`/question`连接地址

**GET `/question?qestion=问题` (Find Question)**

返回示例：

```json
CreatedAt: "0001-01-01T00:00:00Z"
DeletedAt: null
ID: 4
UpdatedAt: "0001-01-01T00:00:00Z"
answer: "XXX"
depart: "XXX"
name: "XXX"
question: "xxx"
```

**POST `/question`(Create Question)**

输入示例

```json
"question":"xxx"
"answer":"xxx"
"depart":"xxx"
"name":"xxx"
```

返回示例：

成功

```json
"status":200,
"message":"Question insert ok",
"message_id":xxx,
```

失败

```json
"message" : "database error"
或者
"message": "error payload"
```

**PUT`/question:id`(update question)**

输入示例

```json
"question":"xxx"
"answer":"xxx"
"depart":"xxx"
"name":"xxx"
"id":id,
```

返回示例：

```json
"status":426,
"message":"Question update ok",
"message_id":ID,
```

失败

```json
"message" : "update error"
或者
"message": "error payload"
```

**DELETE`/question:id`**

```json
"message":"Delete Error",
```

```json
"message":"Delete Successful",
```



**感谢E99p1ant小哥哥的倾力支持**

