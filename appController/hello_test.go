package appController

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloValid(t *testing.T) {
	e, _ := initTestEcho()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/")
	err := Hello(c)
	if err != nil {
		t.Errorf("should not get error, get error: %s", err)
		return
	}
	if rec.Code != 200 {
		t.Errorf("should return 200, get: %d", rec.Code)
	}
	actual := rec.Body.String()
	expected := "hello world 🌏"
	if actual != expected {
		t.Errorf("should return %s, get: %s", expected, actual)
	}
}
