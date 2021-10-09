package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "forms.html")
	case "POST":

		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
		name := r.FormValue("name")
		email := r.FormValue("email")
		id := r.FormValue("id")
		pass := r.FormValue("password")
		fmt.Fprintf(w, "Name = %s\n", name)
		fmt.Fprintf(w, "email = %s\n", email)
		fmt.Fprintf(w, "Username = %s\n", id)
		fmt.Fprintf(w, "Password = %s\n", pass)
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func main() {
	http.HandleFunc("/", hello)

	fmt.Printf("Starting ...\n")
	if err := http.ListenAndServe(":5500", nil); err != nil {
		log.Fatal(err)
	}

}
