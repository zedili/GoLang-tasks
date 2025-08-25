package api

import (
	"main/helper"
	"main/model"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	BaseController
}

func (ctl *UserController) Router(rg *gin.RouterGroup) {
	rg.POST("/login", ctl.Login)
	rg.POST("register", ctl.Register)
}

func (ctl *UserController) Register(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		ctl.Fail(c, "注册失败", err)
		return
	}
	user.Password = string(hashedPassword)

	if err := helper.GerMyqlDb().Create(&user).Error; err != nil {
		ctl.Fail(c, "注册失败", err)
		return
	}

	ctl.Success(c, "注册成功", nil)

}

func (ctl *UserController) Login(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		// c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctl.Fail(c, "登陆失败:", err)
		return
	}

	var storedUser model.User
	if err := helper.GerMyqlDb().Where("username = ?", user.Username).First(&storedUser).Error; err != nil {
		// c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		ctl.Fail(c, "用户名或者密码错误", err)
		return
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)); err != nil {
		// c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		ctl.Fail(c, "用户名或者密码错误", err)
		return
	}

	// 生成 JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       storedUser.ID,
		"username": storedUser.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte("your_secret_key"))
	if err != nil {
		// c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		ctl.Fail(c, "登录失败", err)
		return
	}

	ctl.Success(c, "登陆成功", tokenString)

}
