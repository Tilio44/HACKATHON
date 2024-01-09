// main.go
package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

var (
	Posts []Post
)

type Post struct {
	Author  string `json:"author"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Date    string `json:"date"`
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

func server() {
	loadPostsFromFile()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			displayPosts(w)
		} else if r.Method == "POST" {
			addPost(w, r)
		}
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Erreur lors du démarrage du serveur :", err)
	}
}

func displayPosts(w http.ResponseWriter) {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Println("Erreur lors du rendu du template HTML :", err)
		return
	}

	err = tmpl.Execute(w, struct{ Posts []Post }{Posts})
	if err != nil {
		fmt.Println("Erreur lors de l'exécution du template HTML :", err)
		return
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

func main() {
	server()

    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	savePostsToFile()
}
