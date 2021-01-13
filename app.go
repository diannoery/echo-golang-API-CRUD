package main

import (
	"echo/databases"
	"echo/routes"
)

func main() {
	// e := echo.New()
	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, World!")
	// })
	// e.GET("/users/:id", getUser)
	// e.GET("/show", getShow)
	// // e.POST("/users", saveUser)

	// // e.PUT("/users/:id", updateUser)
	// // e.DELETE("/users/:id", deleteUser)

	databases.Init()
	e := routes.Init()
	e.Logger.Fatal(e.Start(":1323"))

}

// // e.GET("/users/:id", getUser)
// func getUser(c echo.Context) error {
// 	// User ID from path `users/:id`
// 	id := c.Param("id")
// 	return c.String(http.StatusOK, id)
// }

// func getShow(c echo.Context) error {
// 	return c.String(http.StatusOK, "okey")
// }
