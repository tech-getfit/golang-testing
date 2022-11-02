package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestRouter(t *testing.T) {
	router := httprouter.New()

	router.GET("/", func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(rw, "Hello Get")
	})

	request := httptest.NewRequest("GET", "http://localhost:3000/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	byte, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Hello Get", string(byte))
}

func TestRouterParams(t *testing.T) {
	router := httprouter.New()

	router.GET("/item/:id", func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
		text := "Item " + p.ByName("id")

		fmt.Fprint(rw, text)
	})

	request := httptest.NewRequest("GET", "http://localhost:3000/item/1", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	byte, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Item 1", string(byte))
}
