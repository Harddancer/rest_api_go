package apiserver

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestApiserverHandleHello(t *testing.T){

	s  := New(Newconfig())
	rec := httptest.NewRecorder()
	fmt.Println(rec)
	req, _ := http.NewRequest(http.MethodGet,"/hello",nil)
	fmt.Println(req)
	s.handleHello().ServeHTTP(rec,req)
	assert.Equal(t,rec.Body.String(),"Hello")
	


}