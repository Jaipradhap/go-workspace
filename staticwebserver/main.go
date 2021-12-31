package main

import (
	"fmt"
	"log"
	"net/http"
)

type msghandler struct {
	message string
}

func (m *msghandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, m.message)
}

func main() {

	// Static Web Server - File Server
	//1- create multiplexor
	// mux := http.NewServeMux()

	// //2- create handler & assign to mux
	// fs := http.FileServer(http.Dir("public"))
	// mux.Handle("/", fs)

	// //3- tcp port listerner -- func ListenAndServe(addr string, handler Handler) error
	// http.ListenAndServe(":8088", mux)

	// Custom Handlers

	mux1 := http.NewServeMux()

	ch1 := &msghandler{"Hello custom handler"}
	mux1.Handle("/msg", ch1)

	log.Println("Listening 8089")
	http.ListenAndServe(":8089", mux1)

}
