package routes

import (
	"echo/controller"
	"echo/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/pegawai", controller.FecthAll, middleware.IsAuthenticated)
	e.POST("/pegawai", controller.StorePegawai)
	e.PUT("/pegawai", controller.UpdatePegawai)
	e.DELETE("/pegawai", controller.DeletePegawai)

	e.GET("/generet-hash/:password", controller.HashPassword)
	e.POST("/login", controller.CheckPass)

	return e
}
