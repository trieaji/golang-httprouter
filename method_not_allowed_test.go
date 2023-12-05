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
Method Not Allowed Handler -> digunakan jika client tidak mengirim mengirim method sesuai dengan yg kita tentukan
*/

func TestMethodNotAllowed(t *testing.T) {
	router := httprouter.New()
	router.MethodNotAllowed = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request){
		fmt.Fprint(writer, "Gak Boleh lhoo")
	})

	router.POST("/", func (writer http.ResponseWriter, request *http.Request, params httprouter.Params)  { //method post
		fmt.Fprint(writer, "Hello POST")
	})

	request := httptest.NewRequest("GET", "http://localhost:4000/",nil) //nil itu body //mengirim method menggunakan GET
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder,request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Gak Boleh lhoo", string(body))
}
