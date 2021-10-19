package appController

import (
	"fmt"
	"net/http"
	"strconv"

	"gofrendi/structureExample/appMiddleware"
	"gofrendi/structureExample/appModel"

	"github.com/labstack/echo/v4"
)

type LoginInfo struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type PersonController struct {
	model     appModel.PersonModel
	jwtSecret string
}

func NewPersonController(jwtSecret string, m appModel.PersonModel) PersonController {
	return PersonController{
		jwtSecret: jwtSecret,
		model:     m,
	}
}

func (pc PersonController) Login(c echo.Context) error {
	loginInfo := LoginInfo{}
	c.Bind(&loginInfo)
	person, err := pc.model.GetByEmailAndPassword(loginInfo.Email, loginInfo.Password)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "cannot login")
	}
	token, err := appMiddleware.CreateToken(int(person.ID), pc.jwtSecret)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "cannot login")
	}
	person.Token = token
	person, err = pc.model.Edit(int(person.ID), person)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "cannot add token")
	}
	return c.JSON(http.StatusOK, person)
}

func (pc PersonController) GetAll(c echo.Context) error {
	currentLoginPersonId := appMiddleware.ExtractTokenUserId(c)
	fmt.Println("😸 Current user id: ", currentLoginPersonId)
	allPersons, err := pc.model.GetAll()
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "cannot get persons")
	}
	return c.JSON(http.StatusOK, allPersons)
}

func (pc PersonController) Add(c echo.Context) error {
	var person appModel.Person
	if err := c.Bind(&person); err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "invalid person data")
	}
	person, err := pc.model.Add(person)
	if err != nil {
		return c.String(http.StatusInternalServerError, "cannot add person")
	}
	return c.JSON(http.StatusOK, person)
}

func (pc PersonController) Edit(c echo.Context) error {
	personId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "invalid person id")
	}
	var person appModel.Person
	if err := c.Bind(&person); err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "invalid person data")
	}
	person, err = pc.model.Edit(personId, person)
	if err != nil {
		return c.String(http.StatusInternalServerError, "cannot edit person")
	}
	return c.JSON(http.StatusOK, person)
}
