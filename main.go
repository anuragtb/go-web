package main

import (
	"github.com/labstack/echo"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/starptech/go-web/config"
	"github.com/starptech/go-web/controllers"
	"github.com/starptech/go-web/core"
	"github.com/starptech/go-web/models"
)

func main() {
	config := config.NewConfig()

	// create server
	server := core.NewServer(config)
	// serve files for dev
	server.ServeStaticFiles()

	userCtrl := &controllers.User{Context: server}
	feedCtrl := &controllers.Feed{Context: server}
	healthCtrl := &controllers.Healthcheck{Context: server}
	importCtrl := &controllers.Importer{Context: server}

	// api rest endpoints
	g := server.Echo.Group("/api")
	g.GET("/users/:id", userCtrl.GetUserJSON)

	// pages
	u := server.Echo.Group("/users")
	u.GET("/:id", userCtrl.GetUser)
	u.GET("/:id/details", userCtrl.GetUserDetails)

	// special endpoints
	server.Echo.POST("/import", importCtrl.ImportUser)
	server.Echo.GET("/feed", feedCtrl.GetFeed)

	// metric / health endpoint according to RFC 5785
	server.Echo.GET("/.well-known/health-check", healthCtrl.GetHealthcheck)
	server.Echo.GET("/.well-known/metrics", echo.WrapHandler(promhttp.Handler()))

	// migration for dev
	user := models.User{Name: "peter"}
	err := server.GetDB().Register(user)
	if err != nil {
		server.Echo.Logger.Fatal(err)
	}
	server.GetDB().AutoMigrateAll()
	server.GetDB().Create(&user)

	// listen
	go func() {
		server.Echo.Logger.Fatal(server.Start(config.Address))
	}()

	server.GracefulShutdown()
}
