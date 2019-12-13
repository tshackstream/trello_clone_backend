package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/leekchan/timeutil"
	"net/http"
	"os"
	"time"
	"trello_clone_backend/endpoints/borads"
)

func main() {
	e := echo.New()
	t := time.Now()
	fp, err := os.OpenFile("logs/"+timeutil.Strftime(&t, "%Y-%m-%d")+".log",
		os.O_RDWR|os.O_CREATE|os.O_APPEND,
		0666)
	if err != nil {
		panic(err)
	}

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: fp,
	}))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{os.Getenv("ALLOW_ORIGIN")},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost},
	}))

	routing(e)
	port := os.Getenv("ECHO_PORT")
	e.Logger.Fatal(e.Start(":" + port))
}

func routing(e *echo.Echo) {
	e.GET("/api/v1/boards", borads.List)
}
