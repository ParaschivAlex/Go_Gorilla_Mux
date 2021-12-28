package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter() // router object

	func1 := func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		w.Write([]byte(vars["id"]))
	}
	/*func2 := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("function 2"))
	}*/
	//r.PathPrefix("/foo").HandlerFunc(func1)
	//r.PathPrefix("/bar").HandlerFunc(func2)

	r.HandleFunc("/foo/{id:[0-9]+}", func1)
	//r.HandleFunc("/", func1).Schemes("http")
	//r.HandleFunc("/", func2).Schemes("https")

	http.Handle("/", r)

	go http.ListenAndServe(":4000", nil)
	//go http.ListenAndServeTLS(":4443", "cert.pem", "key.pem", nil)

	fmt.Scanln()
}
