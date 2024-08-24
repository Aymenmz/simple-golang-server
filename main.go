package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request){

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() error %s", err)
		return
	}

	fmt.Fprintf(w, "Post Method Succed \n")
	name := r.FormValue("Name")
	address := r.FormValue("Address")

	fmt.Fprintf(w, "Name :: %s \n", name)
	fmt.Fprintf(w, "Adress :: %s \n", address)

	

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 NOT FOUND", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method not Supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello World from Aymen")
}
func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("The Server is Starting on port 8090")

	if err := http.ListenAndServe(":8090", nil); err != nil {
		log.Fatal(err)
	}
}