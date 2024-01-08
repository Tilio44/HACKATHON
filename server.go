// main.go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
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
	fmt.Fprintf(w, `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Posts</title>
    <link rel="stylesheet" href="post.css">
</head>
<body>
    <h1>Liste des posts:</h1>
`)

	for _, post := range Posts {
		fmt.Fprintf(w, `
    <div class="post">
        <h2>%s</h2>
        <p>%s</p>
        <small>Author: %s | Date: %s</small>
    </div>
`, post.Title, post.Content, post.Author, post.Date)
	}

	fmt.Fprintf(w, `
</body>
</html>
`)
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

func generateRandomPosts(numPosts int) {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < numPosts; i++ {
		author := generateRandomString(8)
		title := generateRandomString(10)
		content := generateRandomString(50)
		date := time.Now().Format("2006-01-02")

		newPost := Post{
			Author:  author,
			Title:   title,
			Content: content,
			Date:    date,
		}

		Posts = append(Posts, newPost)
	}
}

func generateRandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}

func main() {
	server()

	// Générer 5 posts aléatoires au démarrage du serveur
	generateRandomPosts(5)

	// Sauvegarder les posts générés dans le fichier JSON
	savePostsToFile()
}
