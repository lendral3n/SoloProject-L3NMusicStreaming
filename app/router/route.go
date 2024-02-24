package router

import (
	ud "l3nmusic/features/user/data"
	uh "l3nmusic/features/user/handler"
	us "l3nmusic/features/user/service"
	"l3nmusic/utils/cloudinary"
	"l3nmusic/utils/encrypts"
	"l3nmusic/utils/middlewares"

	"l3nmusic/app/cache"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo, rds cache.Redis) {
	hash := encrypts.New()
	cloudinaryUploader := cloudinary.New()

	userData := ud.New(db)
	userService := us.New(userData, hash)
	userHandlerAPI := uh.New(userService, cloudinaryUploader)

	// define routes/ endpoint USER
	e.POST("/login", userHandlerAPI.Login)
	e.POST("/users", userHandlerAPI.RegisterUser)
	e.GET("/users", userHandlerAPI.GetUser, middlewares.JWTMiddleware())
	e.PUT("/users", userHandlerAPI.UpdateUser, middlewares.JWTMiddleware())
	e.DELETE("/users", userHandlerAPI.DeleteUser, middlewares.JWTMiddleware())
	e.PUT("/change-password", userHandlerAPI.ChangePassword, middlewares.JWTMiddleware())

}
