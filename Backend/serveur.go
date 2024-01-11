package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

// Post struct represents a blog post
type Post struct {
	Author  string `json:"author"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Date    string `json:"date"`
}

var Posts []Post

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
	renderTemplate(w, "homeActu.page.html", struct{ Posts []Post }{Posts})
}

func main() {
	fmt.Println("http://localhost:8080")
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("../Front/assets"))))
	loadPostsFromFile()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			homeHandler(w, r)
		} else if r.Method == "POST" {
			addPost(w, r)
		}
	})
	http.HandleFunc("/actu", actuHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Erreur lors du démarrage du serveur :", err)
	}
	savePostsToFile()
}

func createTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("../Front/page/*.page.html")
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

func loadPostsFromFile() {
	data, err := ioutil.ReadFile("posts.json")
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier JSON :", err)
		return
	}
	err = json.Unmarshal(data, &Posts)
	if err != nil {
		fmt.Println("Erreur lors de la désérialisation du fichier JSON :", err)
	}
}

func savePostsToFile() {
	data, err := json.Marshal(Posts)
	if err != nil {
		fmt.Println("Erreur lors de la sérialisation en JSON :", err)
		return
	}
	err = ioutil.WriteFile("posts.json", data, 0644)
	if err != nil {
		fmt.Println("Erreur lors de l'écriture dans le fichier JSON :", err)
	}
}

func addPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	author := r.Form.Get("author")
	title := r.Form.Get("title")
	content := r.Form.Get("content")
	date := r.Form.Get("date")

	newPost := Post{
		Author:  author,
		Title:   title,
		Content: content,
		Date:    date,
	}

	Posts = append(Posts, newPost)
	savePostsToFile()

	fmt.Fprintf(w, "Post ajouté avec succès:\nTitle: %s\nContent: %s\nDate: %s\n", title, content, date)
}
