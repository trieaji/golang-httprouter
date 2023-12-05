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

/* 
1. Selain panic handler, Router juga memiliki not found handler
2. Not found handler adalah handler yg dieksekusi ketika client mencoba melakukan request URL yg memang tidak terdapat di router
3. Secara default, jika tidak ada route tidak ditemukan, Router akan melanjutkan request ke http.NotFound, namun kita bisa mengubahnya
4. Caranya dengan mengubah router.NotFound = http.Handler
*/

func TestNotFound(t *testing.T) {
	router := httprouter.New()
	
	router.NotFound = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request){//Caranya dengan mengubah router.NotFound = http.Handler
		fmt.Fprint(writer, "Gak Ketemu")
	})

	request := httptest.NewRequest("GET", "http://localhost:4000/",nil) //nil itu body
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder,request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Gak Ketemu", string(body))
}
