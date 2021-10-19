package appController

import (
	"fmt"
	"gofrendi/structureExample/arithmetic"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func Add(c echo.Context) error {
	firstNum, err := strconv.Atoi(c.Param("firstNum"))
	if err != nil {
		return c.String(http.StatusBadRequest, "firstNum invalid")
	}
	secondNum, err := strconv.Atoi(c.Param("secondNum"))
	if err != nil {
		return c.String(http.StatusBadRequest, "secondNum invalid")
	}
	result := fmt.Sprintf("%d", arithmetic.Add(firstNum, secondNum))
	return c.String(http.StatusOK, result)
}
