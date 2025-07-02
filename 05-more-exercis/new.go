package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		fmt.Fprintln(w, "Hello, World!")
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "About Page")
	})

	http.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			w.Header().Set("Content-Type", "text/html")
			fmt.Fprintln(w, `<form method="POST" action="/submit">`)
			fmt.Fprintln(w, `Name: <input name="name" type="text" /> <input type="submit" value="Submit" />`)
			fmt.Fprintln(w, `</form>`)
			return
		}
		if r.Method == http.MethodPost {
			name := r.FormValue("name")
			if name == "" {
				fmt.Fprintln(w, "Please provide a name.")
				return
			}
			fmt.Fprintf(w, "Hello, %s!", name)
			return
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(w, "Method not allowed.")
	})

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
