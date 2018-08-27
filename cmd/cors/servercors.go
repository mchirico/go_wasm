package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", RootEndpoint)
	r.HandleFunc("/users", UserEndpoint)
	r.HandleFunc("/projects", ProjectEndpoint)

	// Apply the CORS middleware to our top-level router, with the defaults.
	log.Printf("Error: %v", http.ListenAndServe(":8000", handlers.CORS()(r)) )
}

func RootEndpoint(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	//vars["title"] // the book title slug
	//vars["page"] // the page

	str := fmt.Sprintf(`/users
/projects
Root: vars: %v\n`, vars)

	w.Write([]byte(str))

}

func UserEndpoint(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	//vars["title"] // the book title slug
	//vars["page"] // the page
	log.Printf("User vars: %v\n", vars)
	w.Write([]byte("/UserEndpoint"))

}

func ProjectEndpoint(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	//vars["title"] // the book title slug
	//vars["page"] // the page
	log.Printf("vars: %v\n", vars)
	w.Write([]byte("/ProjectEndpoint"))
}
