package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenEverytingOk(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	status := responseRecorder.Code

	require.Equal(t, status, http.StatusOK)

	body := responseRecorder.Body.String()
	assert.NotEmpty(t, body)
	assert.Equal(t, body, "Мир кофе,Сладкоежка")
}

func TestWhenWrongCity(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=paris", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	status := responseRecorder.Code

	require.Equal(t, status, http.StatusBadRequest)
	assert.Equal(t, responseRecorder.Body.String(), "wrong city value")
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	status := responseRecorder.Code

	require.Equal(t, status, http.StatusOK)

	body := responseRecorder.Body
	assert.NotEmpty(t, body)

	res := strings.Split(body.String(), ",")
	assert.Equal(t, len(res), totalCount)
	assert.Equal(t, res, []string{"Мир кофе", "Сладкоежка", "Кофе и завтраки", "Сытый студент"})
}
