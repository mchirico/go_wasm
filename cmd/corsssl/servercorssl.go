package main

import (
	"crypto/tls"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", RootEndpoint)
	r.HandleFunc("/users", UserEndpoint)
	r.HandleFunc("/projects", ProjectEndpoint)

	cfg := &tls.Config{
		MinVersion:               tls.VersionTLS12,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		},
	}
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      handlers.CORS()(r),
		TLSConfig:    cfg,
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
	}
	log.Fatal(srv.ListenAndServeTLS("/etc/certs/fullchain.pem", "/etc/certs/key.pem"))

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
