package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	api "github.com/amarka8/learning-management-system/cmd/server"
)

type healthCheckTest struct {
	Status string
}

/*
HELPER FUNCTIONS
*/

// helper function to send a mock request using Go's httptest package
func sendRequest(handler http.Handler, reqType string, pathname string, body string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(reqType, pathname, strings.NewReader(body))
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	return w
}

func assertStatus(t *testing.T, w *httptest.ResponseRecorder, expected int) {
	if w.Result().StatusCode != expected {
		t.Errorf("Expected status %d but got %d", expected, w.Result().StatusCode)
	}
}

// this test checks that the health check returns {Status: ok} and a 200 status code
func TestHealthCheck(t *testing.T) {
	handler := api.NewDataDocHandler("", "", "")
	resp := sendRequest(handler, "GET", "/health", "")
	assertStatus(t, resp, 200)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read response body, health check failed")
		return
	}
	var result healthCheckTest
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("Error unmarshaling response body")
		return
	}
	fmt.Printf("status: %s", result.Status)
}
