package main

import (
	"net/http"
	"todo/src/auth"
	"todo/src/todo"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	//Logger Middleware
	e.Use(middleware.Logger())

	//Recover Middleware
	e.Use(middleware.Recover())

	e.Use(middleware.CORS())

	//Login
	e.POST("/login", auth.Login)

	//SignUp
	e.POST("/signup", auth.SignUp)

	//root -> 모든요처에 JWT Token
	root := e.Group("")
	{
		//JWT Token Middleware
		config := middleware.JWTConfig{
			Claims:     &auth.JwtClaims{},
			SigningKey: []byte("secret"),
		}
		root.Use(middleware.JWTWithConfig(config))
		root.Use(auth.Authorization)

		root.GET("/test", func(c echo.Context) error {
			return c.String(http.StatusOK, "Hello "+c.QueryParam("userid"))
		})

		root.GET("/todo", todo.QueryTodos)

		root.PUT("/todo", todo.CreateTodos)

		root.PATCH("/todo", todo.UpdateTodos)

		root.DELETE("/todo", todo.DeleteTodos)
	}

	e.Logger.Fatal(e.Start(":1323"))
}
