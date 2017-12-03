package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/context"

	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	r, _ := http.NewRequest("GET", "/hello", nil)
	w := httptest.NewRecorder()

	//Hack to try to fake gorilla/mux vars
	vars := map[string]string{
		"mystring": "abcd",
	}
	context.Set(r, 0, vars)

	hello(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	log.Println(string(w.Body.Bytes()))

}
