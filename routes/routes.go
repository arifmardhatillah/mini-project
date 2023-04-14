package routes

import (
	"prak/constants"
	"prak/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	jwtMiddleware := middleware.JWT([]byte(constants.SECRET_JWT))
	// Route / to handler function
	// Users
	e.GET("/users", controllers.GetUsersController, jwtMiddleware)
	e.GET("/users/:id", controllers.GetUserController, jwtMiddleware)
	e.POST("/users", controllers.CreateUserController)
	e.DELETE("/users/:id", controllers.DeleteUserController, jwtMiddleware)
	e.PUT("/users/:id", controllers.UpdateUserController, jwtMiddleware)
	e.POST("/users/login", controllers.LoginUserController)

	//Books
	e.GET("/books", controllers.GetBooksController, jwtMiddleware)
	e.GET("/books/:id", controllers.GetBookController, jwtMiddleware)
	e.POST("/books", controllers.CreateBookController, jwtMiddleware)
	e.DELETE("/books/:id", controllers.DeleteBookController, jwtMiddleware)
	e.PUT("/books/:id", controllers.UpdateBookController, jwtMiddleware)

	//Blogs
	e.GET("/blogs", controllers.GetBlogsController)
	e.GET("/blogs/:id", controllers.GetBlogController)
	e.POST("/blogs", controllers.CreateBlogController)
	e.DELETE("/blogs/:id", controllers.DeleteBlogController)
	e.PUT("/blogs/:id", controllers.UpdateBlogController)
	return e
}
