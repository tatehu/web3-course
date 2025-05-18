package routes

import (
	"github.com/gin-gonic/gin"
	v1 "personal_blog/api/v1"
	"personal_blog/middleware"
	"personal_blog/utils"
)

func InitRouter() {

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.Logger())
	//设置中间件，以下操作需要用户权限（凭证）
	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{
		//实现文章的创建功能，只有已认证的用户才能创建文章，创建文章时需要提供文章的标题和内容。,
		auth.POST("post/add", v1.CreatePost)
		//实现文章的读取功能，支持获取所有文章列表和单个文章的详细信息。,
		auth.GET("post/:id", v1.GetPost)
		auth.GET("post/all", v1.GetPosts)
		//实现文章的更新功能，只有文章的作者才能更新自己的文章。,
		auth.POST("post/update/:id", v1.UpdatePost)
		//实现文章的删除功能，只有文章的作者才能删除自己的文章。
		auth.DELETE("post/:id", v1.DeletePost)

		//实现评论的创建功能，已认证的用户可以对文章发表评论。,
		auth.POST("comment/add/:id", v1.CreateComment)
		//实现评论的读取功能，支持获取某篇文章的所有评论列表。
		auth.GET("comment/:id", v1.GetComments)
	}
	routeV1 := r.Group("api/v1")
	{

		//登录
		routeV1.POST("login", v1.Login)
		//注册用户
		routeV1.POST("/register", v1.Register)

	}

	r.Run(utils.HttpPort)
}
