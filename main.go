package main

import (
	"hangmanweb/hangman"
	"html/template"
	"net/http"
)

type Infos struct {
	Name string
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/menu", http.StatusFound)
}

func menuHandler(w http.ResponseWriter, r *http.Request) {
	infos := Infos{
		Name: "",
	}

	// Mise à jour du chemin du fichier HTML pour le menu
	tmpl, err := template.ParseFiles("./template/menu.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, infos)
}

func commencerHandler(w http.ResponseWriter, r *http.Request) {
	// Mise à jour du chemin du fichier HTML pour commencer
	go hangman.Init()
	tmpl, err := template.ParseFiles("./template/commencer.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func main() {
	// Rediriger la racine vers /menu
	http.HandleFunc("/", rootHandler)

	// Servir les fichiers CSS depuis le dossier styles
	fsCSS := http.FileServer(http.Dir("./styles"))
	http.Handle("/styles/", http.StripPrefix("/styles/", fsCSS))

	// Servir les images
	fsImg := http.FileServer(http.Dir("./img"))
	http.Handle("/img/", http.StripPrefix("/img/", fsImg))

	// Route pour le menu principal
	http.HandleFunc("/menu", menuHandler)

	// Route pour la page commencer
	http.HandleFunc("/commencer", commencerHandler)

	// Lancer le serveur
	http.ListenAndServe(":1551", nil)
}
