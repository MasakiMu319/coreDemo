package main

import (
	"github.com/luoshengyue/coreDemo/framework/gin"
	"github.com/luoshengyue/coreDemo/framework/middleware"
)

func registerRouter(core *gin.Engine) {
	core.GET("/user/login", middleware.Test3(), UserLoginController)

	subjectApi := core.Group("/subject")
	{
		subjectApi.DELETE("/:id", SubjectDelController)
		subjectApi.PUT("/:id", SubjectUpdateController)
		subjectApi.GET("/:id", middleware.Test3(), SubjectGetController)
		subjectApi.GET("/list/all", SubjectListController)

		subjectInnerApi := subjectApi.Group("/info")
		{
			subjectInnerApi.GET("/name", SubjectNameController)
		}
	}
}
