package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainHandlerWhenRequestIsOK(t *testing.T) {
	params := url.Values{}
	params.Add("count", "4")
	params.Add("city", "moscow")

	req := httptest.NewRequest("GET", "/cafe?"+params.Encode(), nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.NotEmpty(t, responseRecorder.Body)
}

func TestMainHandlerWhenWrongCityValue(t *testing.T) {
	params := url.Values{}
	params.Add("count", "4")
	params.Add("city", "abrbbab")

	req := httptest.NewRequest("GET", "/cafe?"+params.Encode(), nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	assert.Equal(t, "wrong city value", responseRecorder.Body.String())
}

func TestMainHandlerWhenCountMoreThanMaxLen(t *testing.T) {
	params := url.Values{}
	params.Add("count", "10")
	params.Add("city", "moscow")

	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?"+params.Encode(), nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	body := responseRecorder.Body.String()
	list := strings.Split(body, ",")

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.Equal(t, totalCount, len(list))

}
