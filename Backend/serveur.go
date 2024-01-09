package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, "home.html")
}

func homeCSSHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/home.css")
}

func actuHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/actu.html")
}

func actuCSSHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/actu.css")
}

func main() {
	http.HandleFunc("/home", homeHandler)
	http.HandleFunc("/home.css", homeCSSHandler)
	http.HandleFunc("/actu", actuHandler)
	http.HandleFunc("/actu.css", actuCSSHandler)

	fmt.Println("Serveur démarré sur le port : 8080")
	http.ListenAndServe(":8080", nil)
}
