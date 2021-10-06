package routes

import (
	"part2/consts"
	c "part2/controller"
	m "part2/model"
	u "part2/util"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	// user routing
	db := u.InitDB()
	mdl := m.NewUserModel(db)
	bc := c.NewController(consts.JWT_SECRET, mdl)

	e.POST("/login", bc.Login)
	e.GET("/books", bc.GetUsersController)
	e.GET("/books/:id", bc.GetUserController)

	eAuth := e.Group("")
	eAuth.Use(middleware.JWT([]byte(consts.JWT_SECRET)))
	eAuth.POST("/books", bc.CreateUserController)
	eAuth.DELETE("/books/:id", bc.DeleteUserController)
	eAuth.PUT("/books/:id", bc.UpdateUserController)

	return e
}
