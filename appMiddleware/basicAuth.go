package appMiddleware

import (
	"gofrendi/structureExample/appModel"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func DummyBasicAuth(username, password string, c echo.Context) (bool, error) {
	if username == "admin" && password == "admin" {
		return true, nil
	}
	return false, nil
}

func MakePersonBasicAuth(m appModel.PersonModel) middleware.BasicAuthValidator {
	return func(username, password string, c echo.Context) (bool, error) {
		_, err := m.GetByEmailAndPassword(username, password)
		if err == nil {
			return true, nil
		}
		return false, nil
	}
}
