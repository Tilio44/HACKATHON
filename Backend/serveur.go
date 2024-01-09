package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

func renderTemplate(w http.ResponseWriter, tmpl string, p interface{}) {
	var myCache, err = createTemplateCache()
	if err != nil {
		fmt.Println(err)
	}

	t, ok := myCache[tmpl]
	if !ok {
		fmt.Println("Could not get template from cache")
	}

	buffer := new(bytes.Buffer)
	t.Execute(buffer, p)
	buffer.WriteTo(w)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.html", nil)
}

func actuHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "homeActu.page.html", nil)
}

func main() {
	fs := http.FileServer(http.Dir("../Front/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	fmt.Println("http://localhost:8080")
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/actu", actuHandler)
	http.ListenAndServe(":8080", nil)
}

func createTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("../Front/Page/*.page.html")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		myCache[name] = ts
	}

	return myCache, nil
}
