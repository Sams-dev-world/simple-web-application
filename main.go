package main

import (
	"crypto/subtle"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const (
	ADMIN_USERNAME = "admin"
	ADMIN_PASSWORD = "admin"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to my awesome site</h1>")
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<p1>To get in touch please send an email to <a href= \"mailto:support@fortress.com\">support@fortress.com</a></p1>")
}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>This is my faq page</h1>")
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>Sorry but we couldn't find the page you were looking for.</h1>")
}
func basicAuth(handler http.HandlerFunc, realm string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		if !ok || subtle.ConstantTimeCompare([]byte(user), []byte(ADMIN_USERNAME)) != 1 || subtle.ConstantTimeCompare([]byte(pass), []byte(ADMIN_PASSWORD)) != 1 {
			w.Header().Set("WWW-Authenticate", `Basic realm="`+realm+`"`)
			w.WriteHeader(401)
			w.Write([]byte("You are unauthorised to access this application"))
			return
		}
		handler(w, r)
	}
}
func main() {
	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(notFound)
	r.HandleFunc("/", basicAuth(home, "Please enter your username and password"))
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faq", faq)
	err := http.ListenAndServe(":3000", r)
	handlers.CompressHandler(r)
	if err != nil {
		log.Fatal("Unexpected error occured in starting up server: ", err)
	}
}
