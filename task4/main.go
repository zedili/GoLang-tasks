package main

import (
	"fmt"
	"main/api"
	"main/helper"
	"main/model"
	"main/mylog"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleWare() gin.HandlerFunc {
	f := func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token == "" {
			ctx.AbortWithStatusJSON(200, gin.H{
				"message": "请登录",
				"code":    "-1",
			})
			return
		}
		jwtToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
			// 验证签名方法
			// (*jwt.SigningMethodHMAC) 类型断言，尝试将 t.methood 转换为 *jwt.SigningMethodHMAC 类型
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				mylog.Error("tonek 验证错误，%v", token)
				return nil, fmt.Errorf("Token验证错误")
			}
			return []byte("your_secret_key"), nil
		})
		if err != nil {
			ctx.AbortWithStatusJSON(200, gin.H{
				"message": "请登录",
				"code":    "-1",
			})
			return
		}
		jwtClaims, ok := jwtToken.Claims.(jwt.MapClaims)

		if !ok {
			ctx.AbortWithStatusJSON(200, gin.H{
				"message": "请登录",
				"code":    "-1",
			})
			return
		}

		id := jwtClaims["id"]
		username := jwtClaims["username"]

		user := &model.User{}
		result := helper.GerMyqlDb().Where("id = ?", id).First(user)
		if result.Error != nil {
			ctx.AbortWithStatusJSON(200, gin.H{
				"message": "请登录",
				"code":    "-1",
			})
			return
		}
		if user.ID == 0 || user.Username != username {
			ctx.AbortWithStatusJSON(200, gin.H{
				"message": "请登录",
				"code":    "-1",
			})
			return
		}
		ctx.Set("user", user)
	}
	return f
}

func initialModel() {
	helper.GerMyqlDb().AutoMigrate(&model.User{}, &model.Post{}, &model.Comment{})

}

func main() {

	mylog.InitLogger()

	initialModel()

	router := gin.Default()
	authGp := router.Group("/")
	new(api.UserController).Router(authGp)

	userGp := router.Group("/user", AuthMiddleWare())
	new(api.PostController).Router(userGp)
	new(api.CommentController).Router(userGp)

	if err := router.Run(":8090"); err != nil {
		mylog.Error("服务启动错误: %v", err)
		// fmt.Println("服务启动错误：", err)
	}

}
