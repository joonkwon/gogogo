package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	// web app can be easily created
	// just by creating a handler or hanlder function
	// here I will use handler function since it is a bit more simple
	http.HandleFunc("/hello", HelloHandleFunc)
	http.HandleFunc("/hello/", HelloNameHandleFunc)

	fmt.Println("starting server")

	// the first argument of ListenAndServe method is address
	// the address is in the format of ip_address:port or hostname:port of network interface
	// if you omit ip_address or hostname, it will listen to all the interfaces
	http.ListenAndServe(":8080", nil)
}

// HelloHandleFunc will output hello world
func HelloHandleFunc(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

// HelloNameHandleFunc will read name from request and append it to the response
func HelloNameHandleFunc(w http.ResponseWriter, req *http.Request) {
	rps := strings.Split(req.RequestURI, "/")
	if len(rps) != 3 {
		return
	}
	fmt.Fprintf(w, "Hello, %s!", rps[2])
}
