package server

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheckHandler(t *testing.T) {

	req := httptest.NewRequest("GET", "/health-check", nil)

	w := httptest.NewRecorder()
	HealthCheckHandler(w, req)

	resp := w.Result()
	if w.Code != http.StatusOK {
		t.Errorf("Status is wrong: %v", w.Code)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	if string(body) != `{"alive": true}` {
		t.Error("Error there ", string(body))
	}
}
