package main

import (
	_ "database/sql"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"notifications/controllers"
	"strings"
)

var headers = "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header, sec-ch-ua, sec-ch-ua-mobile, sec-fetch-dest, sec-fetch-mode, sec-fetch-site"

func handler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "Hello Notifications!")
}

func main() {
	r := mux.NewRouter()

	headersOk := handlers.AllowedHeaders(strings.Split(headers, ", "))
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "PATCH"})

	s := r.PathPrefix("/api").Subrouter()
	controllers.InitNotifications(s)
	r.HandleFunc("/api", handler)

	fmt.Printf("Server starting")
	log.Fatal(http.ListenAndServe(":80", handlers.CORS(originsOk, headersOk, methodsOk)(r)))
}
