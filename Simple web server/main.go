package main

import (
	"fmt"
	"log"
	"net/http"
)

func formhandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parseform() err %v", err)
	}
	fmt.Fprintf(w, "Post request successful!\n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name: %s\n", name)
	fmt.Fprintf(w, "Address: %s\n", address)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Page not found!", http.StatusNotFound)
		return
	}

	if r.Method == "POST" {
		http.Error(w, "Inappropriate method", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello Folks!")
}
func main() {
	fileserver := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileserver)
	http.HandleFunc("/form", formhandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server...")
	if err := http.ListenAndServe(":8080", nil); err != nil { //Creating a server
		log.Fatal(err)
	}

}
