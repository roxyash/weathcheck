package handler

import (
	"net/http"
	"weathcheck/internal/helpers"

	"github.com/gin-gonic/gin"
)


// @Summary weathercheck
// @Tags weathercheck
// @Description weathercheck
// @ID weathercheck
// @Accept  json
// @Produce  json
// @Param        address   query     string  true  "address"  
// @Success 200 {object} types.ResponseWeatherInfo
// @Failure 400,404 {object} types.ErrorResponse
// @Failure 500 {object} types.ErrorResponse
// @Failure default {object} types.ErrorResponse
// @Router /getweatherinfo [get]
func (h *Handler) getWeather(c *gin.Context) {


	address, ok := c.GetQuery("address")
	if !ok {
		helpers.NewErrorResponse(c, http.StatusBadRequest, "invalid address param")
		return
	}

	weatherInfo, err := h.services.WeatherInfoService.GetWeatherInfo(address)
	if err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	
	c.JSON(http.StatusOK, weatherInfo)
}
