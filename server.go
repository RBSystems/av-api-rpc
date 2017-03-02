package main

import (
	"net/http"

	"github.com/byuoitav/authmiddleware"
	"github.com/byuoitav/av-api-rpc/handlers"
	"github.com/jessemillar/health"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	port := ":8100"
	router := echo.New()
	router.Pre(middleware.RemoveTrailingSlash())
	router.Use(middleware.CORS())

	// Use the `secure` routing group to require authentication
	secure := router.Group("", echo.WrapMiddleware(authmiddleware.Authenticate))

	router.GET("/health", echo.WrapHandler(http.HandlerFunc(health.Check)))

	// router.Get("/buildings", handlers.GetAllBuildings, wso2jwt.ValidateJWT())
	secure.GET("/buildings/:building/rooms/:room", handlers.GetRoomByNameAndBuildingHandler)

	// PUT requests
	secure.PUT("/buildings/:building/rooms/:room", handlers.SendRoomCommands)

	server := http.Server{
		Addr:           port,
		MaxHeaderBytes: 1024 * 10,
	}

	router.StartServer(&server)
}
