package handleHttp

import "github.com/labstack/echo/v4"

func RegisterRoutes(handler HttpHandler,e *echo.Echo)  {
	e.POST("/new",handler.Save)
	e.GET("/:surl",handler.Redirect)
}