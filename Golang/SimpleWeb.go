package main

import "fmt"
import "net/http"

func index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "apa kabar!")
}

func about(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "apa about!")
}

func main() {
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "halo!")
    })

	http.HandleFunc("/index", index)
	
	http.HandleFunc("/about",about)

    fmt.Println("starting web server at http://localhost:3000/")
    http.ListenAndServe(":3000", nil)
}