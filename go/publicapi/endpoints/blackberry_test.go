package endpoints

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func TestGetBlackberries(t *testing.T) {
	req, err := http.NewRequest("GET", "/blackberries", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetBlackberries)
	handler.ServeHTTP(rr, req)

	checkResponseCode(t, http.StatusOK, rr.Code)

	expected := `[{"id":1,"name":"Triple Crown"},{"id":2,"name":"Black Satin"}]`
	if rr.Body.String() != expected+"\n" && rr.Body.String() != expected {
		t.Errorf("unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
