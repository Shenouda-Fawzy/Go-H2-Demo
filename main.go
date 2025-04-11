package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello world!")
	fs := http.FileServer(http.Dir("assets"))

	http.Handle("/", fs)
	http.HandleFunc("/hello", SayHello)
	http.HandleFunc("/hook", hook)
	fmt.Println("Serving https://localhost:1234/")
	log.Fatal(http.ListenAndServeTLS(":1234", "localhost.crt", "localhost.key", nil))
}

func SayHello(w http.ResponseWriter, req *http.Request) {
	log.Println("Hi my protocol is: ", req.Proto)
	fmt.Fprintf(w, "Hi i'm the new protocol H2\n")
}

// hook will handle webhook calling
func hook(w http.ResponseWriter, req *http.Request) {
	log.Printf("%s method to %s\n", req.Method, req.URL)
	err := req.ParseForm() // To application/x-www-form-urlencoded
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}

	fmt.Println(req.Form.Encode())
	fmt.Println("--- Beautified ---")
	fmt.Printf("%#v\n", req.Form)
}
