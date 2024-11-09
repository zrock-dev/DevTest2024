package controller

import (
	"PollSystem/internal/domain"
	"PollSystem/internal/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type PollController struct {
    localPollService service.PollService
}

func NewPollCotroller(localPollService service.PollService) PollController {
    return PollController{
    	localPollService: localPollService,
    }
}

func (controller *PollController) CreatePoll(c echo.Context) (error){
    pollRequest := new(domain.PollRequest)
    if err := c.Bind(pollRequest); err != nil {
      return echo.NewHTTPError(http.StatusBadRequest, err)
    }

    controller.localPollService.CreatePoll(*pollRequest)

    return c.JSON(http.StatusOK, pollRequest)
}
