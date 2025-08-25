package api

import (
	"fmt"
	"main/helper"
	"main/model"

	"github.com/gin-gonic/gin"
)

type PostController struct {
	BaseController
}

func (ctl *PostController) Router(rg *gin.RouterGroup) {
	postgp := rg.Group("/post")
	postgp.POST("/create", ctl.CreatePost)
	postgp.POST("/listPost", ctl.listPost)
	postgp.POST("/getPostById", ctl.getPostById)
	postgp.POST("/updateById", ctl.updateById)
	postgp.POST("/deleteById", ctl.deleteById)

}

type PostRequest struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  uint   `json:"userId"`
}

func (ctl *PostController) CreatePost(ctx *gin.Context) {

	userObj, ok := ctl.GetCurrentUser(ctx)
	if !ok {
		ctl.Fail(ctx, "创建文章失败：获取用户信息失败", nil)
	}

	var request PostRequest
	if err := ctx.ShouldBind(&request); err != nil {
		ctl.Fail(ctx, "创建文章失败", err)
		return
	}

	post := model.Post{
		Title:   request.Title,
		Content: request.Content,
		UserID:  userObj.ID,
	}

	fmt.Println("请求参数：", request)
	fmt.Printf("创建文章：%v", post)

	// gorm 需要通过指针来修改原始结构体，特别是在传入后设置自增id 等数据库生成的字段值
	if err := helper.GerMyqlDb().Create(&post).Error; err != nil {
		ctl.Fail(ctx, "创建文章失败", err)
		return
	}

	ctl.Success(ctx, "创建文章成功", nil)
}

func (ctl *PostController) listPost(ctx *gin.Context) {
	var request PostRequest
	ctx.ShouldBind(&request)

	var list []model.Post
	if request.UserID == 0 {
		if err := helper.GerMyqlDb().Find(&list).Error; err != nil {
			fmt.Println("获取文章列表失败：", err)
			ctl.Fail(ctx, "查询失败", err)
			return
		}
		ctl.Success(ctx, "查询成功", list)
		return
	} else {
		if err := helper.GerMyqlDb().Where("user_id = ?", request.UserID).Find(list).Error; err != nil {
			ctl.Fail(ctx, "查询失败", err)
			return
		}
	}
}

func (ctl *PostController) getPostById(ctx *gin.Context) {
	var request PostRequest
	ctx.ShouldBind(&request)

	var post model.Post
	if err := helper.GerMyqlDb().Where("id = ?", request.ID).First(&post).Error; err != nil {
		ctl.Fail(ctx, "查询失败", err)
		return
	}
	ctl.Success(ctx, "查询成功", post)
}

func (ctl *PostController) updateById(ctx *gin.Context) {
	var request PostRequest
	ctx.ShouldBind(&request)

	userObj, ok := ctl.GetCurrentUser(ctx)
	if !ok {
		ctl.Fail(ctx, "修改文章失败：获取用户信息失败", nil)
	}

	if err := helper.GerMyqlDb().Model(&model.User{}).
		Where("id = ? and user_id = ?", request.ID, userObj.ID).
		Update("title", request.Title).
		Update("content", request.Content).Error; err != nil {
		ctl.Fail(ctx, "修改失败", err)
		return
	}
	ctl.Success(ctx, "修改成功", nil)
}

func (ctl *PostController) deleteById(ctx *gin.Context) {
	var request PostRequest
	ctx.ShouldBind(&request)

	userObj, ok := ctl.GetCurrentUser(ctx)
	if !ok {
		ctl.Fail(ctx, "创建文章失败：获取用户信息失败", nil)
	}

	var delCondi model.Post
	if err := helper.GerMyqlDb().
		Model(&model.Post{}).
		Where("id = ? and user_id = ?", request.ID, userObj.ID).
		First(&delCondi).Error; err != nil {
		ctl.Fail(ctx, "文章不存在", err)
		return
	}
	// 传入实体指针，这样钩子函数才能拿到评论的完整信息
	if err := helper.GerMyqlDb().Delete(&delCondi).Error; err != nil {
		ctl.Fail(ctx, "删除文章失败", err)
		return
	}

	ctl.Success(ctx, "删除成功", nil)
}
