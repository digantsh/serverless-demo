package main

import (
	"net/http"

	"github.com/eawsy/aws-lambda-go-net/service/lambda/runtime/net"
	"github.com/eawsy/aws-lambda-go-net/service/lambda/runtime/net/apigatewayproxy"
	"github.com/gorilla/mux"

)


// Handle ... AWS Handler called by Lambda
var Handle apigatewayproxy.Handler



func init() {
  // Handler setup
	ln := net.Listen()
	Handle = apigatewayproxy.New(ln, []string{"image/png"}).Handle

	// MUX routing for the API calls
	// refer api.go for more details on the call
	r := mux.NewRouter()

	// User Authentication and Registration API
	r.Path("/hello").Methods("GET").HandlerFunc(hello)

	go http.Serve(ln, r)
}


func hello(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Hello, World!"))
}

func main() {

}
