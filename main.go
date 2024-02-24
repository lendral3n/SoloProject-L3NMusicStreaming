package main

import (
	"l3nmusic/app/cache"
	"l3nmusic/app/config"
	"l3nmusic/app/database"
	"l3nmusic/app/router"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.InitConfig()
	dbSql := database.InitDBPostgres(cfg)
	cacheRds := cache.InitRedis()

	e := echo.New()
	e.Use(middleware.CORS())
	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))

	router.InitRouter(dbSql, e, cacheRds)
	//start server and port
	e.Logger.Fatal(e.Start(":8000"))
}
