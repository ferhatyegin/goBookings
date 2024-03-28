package main

import (
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {

	//We can run our tests inside the Run() function after setting up all our
	//variables and needed code. os.Exit() will run every function before program exits
	os.Exit(m.Run())
}


type myHandler struct{}

func (mh *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){

}
