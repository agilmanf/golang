package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err!= nil {
		fmt.Fprintf(w, "ParseForm() err:%v", err)
		return
	}

	fmt.Fprintln(w, "POST SUCCESS")
	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintln(w, "Your name is " + name)
	fmt.Fprintln(w, "You live in " + address)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Methon is not Supported", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "hello!")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Server Running on Port 8080")
	if err:= http.ListenAndServe(":8080", nil); err !=nil {
		log.Fatal(err)
	}
}