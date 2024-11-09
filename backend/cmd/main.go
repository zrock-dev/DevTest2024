package main

import (
	"PollSystem/internal/controller"
	"PollSystem/internal/repository"
	"PollSystem/internal/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	repository := repository.NewLocalStorage()
	service := service.NewLocalStorageService(repository)
	pollController := controller.NewPollCotroller(service)

	e := echo.New()
	e.GET("/api/v1/polls", func(c echo.Context) error {
		return pollController.CreatePoll(c)
	})

	e.POST("/api/v1/polls", func(c echo.Context) error {
        pollController.CreatePoll(c)
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/api/v1/polls:id/votes", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
