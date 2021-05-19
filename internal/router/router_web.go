package router

import (
	"github.com/xinliangnote/go-gin-api/internal/web/controller/admin_handler"
	"github.com/xinliangnote/go-gin-api/internal/web/controller/authorized_handler"
	"github.com/xinliangnote/go-gin-api/internal/web/controller/config_handler"
	"github.com/xinliangnote/go-gin-api/internal/web/controller/dashboard_handler"
	"github.com/xinliangnote/go-gin-api/internal/web/controller/gencode_handler"
	"github.com/xinliangnote/go-gin-api/internal/web/controller/index_handler"
	"github.com/xinliangnote/go-gin-api/internal/web/controller/install_handler"
	"github.com/xinliangnote/go-gin-api/internal/web/controller/tool_handler"
)

func setWebRouter(r *resource) {
	installHandler := install_handler.New(r.logger)
	indexHandler := index_handler.New(r.logger, r.db, r.cache)
	dashboardHandler := dashboard_handler.New(r.logger, r.db, r.cache)
	genCodeHandler := gencode_handler.New(r.logger, r.db, r.cache)
	configInfoHandler := config_handler.New(r.logger, r.db, r.cache)
	authorizedHandler := authorized_handler.New(r.logger, r.db, r.cache)
	toolHandler := tool_handler.New(r.logger, r.db, r.cache)
	adminHandler := admin_handler.New(r.logger, r.db, r.cache)

	web := r.mux.Group("", r.middles.DisableLog())
	{
		// Home page
		web.GET("", indexHandler.View())

		// Installation
		web.GET("/install", installHandler.View())
		web.POST("/install/execute", installHandler.Execute())
		web.POST("/install/restart", installHandler.Restart())

		// Dash board
		web.GET("/dashboard", dashboardHandler.View())

		// Configuration information
		web.GET("/config/email", configInfoHandler.EmailView())
		web.GET("/config/code", configInfoHandler.CodeView())

		// Code generation tool
		web.GET("/gormgen", genCodeHandler.GormView())
		web.POST("/gormgen_exec", genCodeHandler.GormExecute())

		web.GET("/handlergen", genCodeHandler.HandlerView())
		web.POST("/handlergen_exec", genCodeHandler.HandlerExecute())

		// Caller
		web.GET("/authorized/list", authorizedHandler.ListView())
		web.GET("/authorized/add", authorizedHandler.AddView())
		web.GET("/authorized/api/:id", authorizedHandler.ApiView())
		web.GET("/authorized/demo", authorizedHandler.DemoView())

		// Administrator
		web.GET("/admin/list", adminHandler.ListView())
		web.GET("/admin/add", adminHandler.AddView())
		web.GET("/admin/modify_password", adminHandler.ModifyPasswordView())
		web.GET("/admin/modify_info", adminHandler.ModifyInfoView())
		web.GET("/login", adminHandler.LoginView())
		web.GET("/admin/menu", adminHandler.MenuView())

		// Toolbox
		web.GET("/tool/hashids", toolHandler.HashIdsView())
		web.GET("/tool/logs", toolHandler.LogsView())
		web.GET("/tool/cache", toolHandler.CacheView())
		web.GET("/tool/data", toolHandler.DataView())
	}
}
