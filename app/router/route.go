package router

import (
	md "l3nmusic/features/music/data"
	mh "l3nmusic/features/music/handler"
	ms "l3nmusic/features/music/service"
	ud "l3nmusic/features/user/data"
	uh "l3nmusic/features/user/handler"
	us "l3nmusic/features/user/service"
	"l3nmusic/utils/encrypts"
	"l3nmusic/utils/middlewares"
	"l3nmusic/utils/upload"

	"l3nmusic/app/cache"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo, rds cache.Redis) {
	hash := encrypts.New()
	s3Uploader := upload.New()

	userData := ud.New(db, rds)
	userService := us.New(userData, hash)
	userHandlerAPI := uh.New(userService, s3Uploader)

	musicData := md.New(db, rds)
	musicService := ms.New(musicData, userService)
	musicHandlerAPI := mh.New(musicService, s3Uploader)

	// define routes/ endpoint USER
	e.POST("/login", userHandlerAPI.Login)
	e.POST("/users", userHandlerAPI.RegisterUser)
	e.GET("/users", userHandlerAPI.GetUser, middlewares.JWTMiddleware())
	e.PUT("/users", userHandlerAPI.UpdateUser, middlewares.JWTMiddleware())
	e.DELETE("/users", userHandlerAPI.DeleteUser, middlewares.JWTMiddleware())
	e.PUT("/change-password", userHandlerAPI.ChangePassword, middlewares.JWTMiddleware())

	// define routes/ endpoint MUSIC
	e.POST("/music", musicHandlerAPI.CreateMusic, middlewares.JWTMiddleware())
	e.GET("/music", musicHandlerAPI.GetAllMusic)
	e.POST("/music/liked/:music_id", musicHandlerAPI.AddLikedSong, middlewares.JWTMiddleware())
}
