package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter() // router object

	func1 := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("function1"))
	}
	func2 := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("function2"))
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		if r.TLS == nil {
			func1(w, r)
		} else {
			func2(w, r)
		}
	}

	r.HandleFunc("/", handler)
	
	//r.HandleFunc("/", func1).Schemes("http")
	//r.HandleFunc("/", func2).Schemes("https")

	http.Handle("/", r)

	go http.ListenAndServe(":4000", nil)
	go http.ListenAndServeTLS(":4443", "cert.pem", "key.pem", nil)

	fmt.Scanln()
}
