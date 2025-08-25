package api

import (
	"main/helper"
	"main/model"
	"main/mylog"

	"github.com/gin-gonic/gin"
)

type CommentController struct {
	BaseController
}

func (ctl *CommentController) Router(rg *gin.RouterGroup) {
	commetnGp := rg.Group("/comment")
	commetnGp.POST("/create", ctl.CreateComment)
	commetnGp.POST("/listComment", ctl.listComment)
	commetnGp.POST("/deleteById", ctl.deleteById)
}

type CommentRequest struct {
	Id      uint   `json:id`
	PostID  uint   `json:"postId"`
	Content string `json:"content"`
}

func (ctl *CommentController) CreateComment(ctx *gin.Context) {

	userObj, ok := ctl.GetCurrentUser(ctx)
	if !ok {
		ctl.Fail(ctx, "评论失败：获取用户信息失败", nil)
	}

	var request CommentRequest
	if err := ctx.ShouldBind(&request); err != nil {
		ctl.Fail(ctx, "评论文章失败", err)
		return
	}

	comment := model.Comment{
		PostID:  request.PostID,
		Content: request.Content,
		UserID:  userObj.ID,
	}

	// gorm 需要通过指针来修改原始结构体，特别是在传入后设置自增id 等数据库生成的字段值
	if err := helper.GerMyqlDb().Create(&comment).Error; err != nil {
		ctl.Fail(ctx, "评论文章失败", err)
		return
	}

	ctl.Success(ctx, "评论文章成功", nil)
}

func (ctl *CommentController) listComment(ctx *gin.Context) {
	var request CommentRequest
	ctx.ShouldBind(&request)

	var list []model.Comment
	if request.PostID == 0 {
		// 返回空列表
		ctl.Success(ctx, "查询成功", list)
		return
	} else {
		if err := helper.GerMyqlDb().
			Where("post_id = ?", request.PostID).
			Find(&list).Error; err != nil {
			mylog.Error("查询失败：%v", err)
			ctl.Fail(ctx, "查询失败", err)
			return
		}
		ctl.Success(ctx, "查询成功", list)
	}
}

func (ctl *CommentController) deleteById(ctx *gin.Context) {
	var request CommentRequest
	ctx.ShouldBind(&request)

	userObj, ok := ctl.GetCurrentUser(ctx)
	if !ok {
		ctl.Fail(ctx, "获取用户信息失败", nil)
	}

	var delCondi model.Comment
	if err := helper.GerMyqlDb().
		Model(&model.Comment{}).
		Where("id = ? and user_id = ?", request.Id, userObj.ID).
		First(&delCondi).Error; err != nil {
		ctl.Fail(ctx, "评论不存在或者不是当前用户的评论", err)
		return
	}
	// 传入实体指针，这样钩子函数才能拿到评论的完整信息
	if err := helper.GerMyqlDb().Delete(&delCondi).Error; err != nil {
		mylog.Info("删除评论失败：%v", delCondi)
		ctl.Fail(ctx, "删除评论失败", err)
		return
	}
	mylog.Info("删除评论成功：%v", delCondi)
	ctl.Success(ctx, "删除成功", nil)
}
