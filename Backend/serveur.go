package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

// Exercise représente un exercice
type Exercise struct {
	Name string `json:"Name"`
}

// Workout représente une séance d'entraînement
type Workout struct {
	Type      string     `json:"Type"`
	Exercises []Exercise `json:"Exercises"`
}

var workouts = map[string][]Exercise{
	"jambes": {Exercise{"Squat"}, Exercise{"Leg Press"}, Exercise{"Lunges"}, Exercise{"Leg Curls"}, Exercise{"Calf Raises"}},
	"pecs":   {Exercise{"Bench Press"}, Exercise{"Incline Press"}, Exercise{"Flyes"}, Exercise{"Push-ups"}, Exercise{"Dips"}},
	"dos":    {Exercise{"Deadlift"}, Exercise{"Pull-ups"}, Exercise{"Bent Over Rows"}, Exercise{"Lat Pulldowns"}, Exercise{"Face Pulls"}},
	// Ajoutez d'autres types de séances avec leurs exercices associés
}

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
	renderTemplate(w, "Hurluberlu.page.html", nil)
}

func workoutHandler(w http.ResponseWriter, r *http.Request) {
	workoutType := r.URL.Path[len("/workout/"):]
	exercises, exists := workouts[workoutType]
	if !exists {
		http.Error(w, "Type de séance de sport non trouvé", http.StatusNotFound)
		return
	}

	workout := Workout{Type: workoutType, Exercises: exercises}
	renderTemplate(w, "workout.tmpl", workout)
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

func main() {
	fmt.Println("http://localhost:8080")
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("../Front/assets"))))

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/workout/", workoutHandler)
	// Ajoutez d'autres gestionnaires pour les différentes pages

	http.ListenAndServe(":8080", nil)
}
