package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name string
	Age  uint
}

/*
*
定义 gin 的中间件函数

执行顺序：
请求到达

	↓

中间件1-before

	↓

中间件2-before

	↓

路由处理函数

	↓

中间件2-after

	↓

中间件1-after

	↓

响应数据

*
*/
func MiddleWare1() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("midware1 before") // 请求到达中间件时执行
		ctx.Next()                     // 调用下一个中间件或接口本身
		fmt.Println("midware1 after")  //
	}
}

func MiddleWare2() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("midware2 before") //
		ctx.Next()                     //
		fmt.Println("midware2 after")  //
	}
}

func MiddleWare3() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("midware3 before") //
		ctx.Next()                     //
		fmt.Println("midware3 after")  //
	}
}

func MiddleWare4() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("midware4 before") // 请求到达中间件时执行
		ctx.Next()                     // 将控制权交给后续的中间件或路由处理函数
		fmt.Println("midware4 after")  //
	}
}

func main() {
	r := gin.Default()
	// 使用全局中间件
	r.Use(MiddleWare1())
	r.Use(MiddleWare2())
	r.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "hello gin111")
	})

	r.GET("/123", func(ctx *gin.Context) {
		ctx.String(200, "hello gin")
	})

	// 接口重定向
	r.GET("/user", MiddleWare3(), func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, "/v1/user")
	})

	// 接口中间件
	r.GET("/md", MiddleWare3(), func(ctx *gin.Context) {
		ctx.String(200, "hello interface middleware")
	})

	r.GET("/baidu", func(ctx *gin.Context) {
		// ctx.Header("Location", "https://www.baidu.com")
		// ctx.AbortWithStatus(http.StatusFound)
		ctx.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
	})

	r.Static("static", "./static")
	r.StaticFS("staticfs", http.Dir("static"))
	r.StaticFile("staticfile", "./static/1.txt")

	r.LoadHTMLGlob("./templates/*")
	r.GET("/index", func(ctx *gin.Context) {
		ctx.HTML(200, "index.impl", gin.H{
			"title": "html 测试",
		})
	})

	// 组接口使用中间件
	gp1 := r.Group("/v1", MiddleWare4())
	{
		gp1.GET("/user", func(ctx *gin.Context) {
			var usre User
			ctx.ShouldBind(&usre)
			ctx.String(200, "v1 user")
		})
	}

	// 基本认证框架
	auth := r.Group("/auth", gin.BasicAuth(gin.Accounts{
		"user1": "123456",
	}))

	auth.GET("/name", func(ctx *gin.Context) {
		user := ctx.MustGet(gin.AuthUserKey).(string)
		ctx.String(http.StatusOK, user)
	})

	err := r.Run(":8080")
	// r.RunTLS()  //使用 https 服务曾
	if err != nil {
		panic(err)
	}

}
