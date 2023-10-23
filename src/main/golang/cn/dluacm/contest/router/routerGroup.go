package router

import (
	"acm-registration-system/middleware/limiter"
	"acm-registration-system/pkg/logger"
	"io"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Init() {
	engine := gin.Default()
	store := cookie.NewStore([]byte("acm-registration-system-cookie"))
	engine.Use(sessions.Sessions("acm-registration-system-session", store))
	engine.ForwardedByClientIP = true
	gin.DefaultWriter = io.MultiWriter(os.Stdout, logger.Logger)
	// 配置跨域 限流
	engine.Use(limiter.Limiter())
	// engine.Use(cors.Cors(), limiter.Limiter())

	// 总路由
	api := engine.Group("/contest/myadmin")
	// 注册用户模块路由
	RegisterUserRouter(api)
	// 注册信息模块路由
	RegisterInfoRouter(api)
	// 注册日志路由
	RegisterLogRouter(api)
	// 注册题目路由
	RegisterFileRouter(api)
	// 注册文件路由
	err := engine.Run(":8183")
	if err != nil {
		logrus.Fatal("服务启动失败")
	}
}
