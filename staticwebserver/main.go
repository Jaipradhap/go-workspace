package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type msghandler struct {
	message string
}

func (m *msghandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, m.message)
}

func main() {

	// %%%%%%%%%%% Static Web Server - File Server
	//1- create multiplexor
	// mux := http.NewServeMux()

	// //2- create handler & assign to mux
	// fs := http.FileServer(http.Dir("public"))
	// mux.Handle("/", fs)

	// //3- tcp port listerner -- func ListenAndServe(addr string, handler Handler) error
	// http.ListenAndServe(":8088", mux)

	// %%%%%%%%%%% Custom Handlers using struct
	/*
		mux1 := http.NewServeMux()

		ch1 := &msghandler{"Hello custom handler"}
		mux1.Handle("/msg", ch1)

		ch2 := &msghandler{"Hello 2"}
		mux1.Handle("/msg2", ch2)

		log.Println("Listening 8089")
		http.ListenAndServe(":8089", mux1)

	*/

	// %%%%%%%%%%% func handlers http.HandlerFunc  & ServeMux.HandlerFunc
	// Step -------------1
	mux2 := http.NewServeMux()

	// httpfunc := http.HandlerFunc(httpfunctionhandler)
	// mux2.Handle("/hfunc", httpfunc)

	// Step -------------2
	// 1- file server
	fs := http.FileServer(http.Dir("public"))
	mux2.Handle("/", fs)

	// 2- custom handler
	ch1 := &msghandler{"Hello custom handler"}
	mux2.Handle("/msg1", ch1)

	// 3- closure logic handler to sent db connection name
	mux2.Handle("/hfunc", closureHandler("calling closureHandler"))

	log.Println("Listening ...")
	// 4- http.Server is used to set server config
	// server := &http.Server{
	// 	Addr:           ":8088",
	// 	ReadTimeout:    10 * time.Second,
	// 	WriteTimeout:   10 * time.Second,
	// 	MaxHeaderBytes: 1 << 20,
	// 	Handler:        mux2,
	// }
	// http.ListenAndServe(":8099", mux2)

	// Step -------------3
	//tcp port listerner -- func ListenAndServe(addr string, handler Handler) error
	// server.ListenAndServe()
	err := listenserverfunction(":8089", mux2)
	if err != nil {
		log.Fatal("Server in error")
	}

	// err1 := listenserverfunction(":8088", mux2)
	// if err1 != nil {
	// 	log.Fatal("Server in error")
	// }
}

//http.HandlerFunc
func httpfunctionhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome func handler")
}

//Handler logic into closure - access outside variable in anonymous fucntion
func closureHandler(msg string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, msg)
	})
}

func listenserverfunction(addr string, h http.Handler) error {
	server := &http.Server{
		Addr:           addr,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
		Handler:        h,
	}

	return server.ListenAndServe()

}
