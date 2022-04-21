package appController

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSubValid(t *testing.T) {
	e, _ := initTestEcho()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/add/3/4")
	c.SetParamNames("firstNum", "secondNum")
	c.SetParamValues("5", "4")
	err := Sub(c)
	if err != nil {
		t.Errorf("should not get error, get error: %s", err)
		return
	}
	if rec.Code != 200 {
		t.Errorf("should return 200, get: %d", rec.Code)
	}
	actual := rec.Body.String()
	expected := "1"
	if actual != expected {
		t.Errorf("should return %s, get: %s", expected, actual)
	}
}

func TestSubInvalidFirstNum(t *testing.T) {
	e, _ := initTestEcho()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/add/tiga/4")
	c.SetParamNames("firstNum", "secondNum")
	c.SetParamValues("tiga", "4")
	err := Sub(c)
	if err != nil {
		t.Errorf("should not get error, get error: %s", err)
		return
	}
	if rec.Code != 400 {
		t.Errorf("should return 400, get: %d", rec.Code)
	}
	actual := rec.Body.String()
	expected := "firstNum invalid"
	if actual != expected {
		t.Errorf("should return %s, get: %s", expected, actual)
	}
}

func TestSubInvalidSecondNum(t *testing.T) {
	e, _ := initTestEcho()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/add/3/empat")
	c.SetParamNames("firstNum", "secondNum")
	c.SetParamValues("3", "empat")
	err := Sub(c)
	if err != nil {
		t.Errorf("should not get error, get error: %s", err)
		return
	}
	if rec.Code != 400 {
		t.Errorf("should return 400, get: %d", rec.Code)
	}
	actual := rec.Body.String()
	expected := "secondNum invalid"
	if actual != expected {
		t.Errorf("should return %s, get: %s", expected, actual)
	}
}
