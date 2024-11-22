package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMessageHandler(t *testing.T) {
	message := map[string]string{"message": "Hello, World!"}
	jsonMessage, _ := json.Marshal(message)
	req, err := http.NewRequest("POST", "/message", bytes.NewBuffer(jsonMessage))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(MessageHandler)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "Expected status code 200")

	expected := `{"status":"success","message":"Hello, World!"}`
	assert.JSONEq(t, expected, rr.Body.String(), "Response body mismatch")
}
