package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./main.html"))
	tmpl.Execute(w, nil)
}

func contact(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./contact.html"))
	tmpl.Execute(w, nil)
}

func about(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./about.html"))
	tmpl.Execute(w, nil)
}

func video(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./video.html"))
	tmpl.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", handler)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./css"))))
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/about", about)
	http.HandleFunc("/video", video)
	port := 8080
	fmt.Printf("Serveur en cours d'exécution sur le port %d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Erreur lors du démarrage du serveur:", err)
	}
}
