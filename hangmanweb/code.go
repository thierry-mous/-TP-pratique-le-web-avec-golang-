package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

var traitement Donnes

func main() {

	temp, err := template.ParseGlob("./template/*.html")
	if err != nil {
		fmt.Println(fmt.Sprintf("Erreur => %s", err.Error()))
		return
	}

	http.HandleFunc("/var", func(w http.ResponseWriter, r *http.Request) {
		Promotion := promo{"Mentor'ac",
			"informatique",
			5,
			3,
			[]Etudiant{
				{"Cyril", "Rodrigues", 22, "homme"},
				{"Kheir-eddine", "Merderreg", 22, "homme"},
				{"Alan", "Philipiert", 26, "homme"}}}
		temp.ExecuteTemplate(w, "var", Promotion)
	})

	http.HandleFunc("/user/init", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "init", nil)
	})

	http.HandleFunc("/user/treatment", func(w http.ResponseWriter, r *http.Request) {
		traitement = Donnes{
			r.FormValue("user_nom"),
			r.FormValue("user_prenom"),
			r.FormValue("user_date"),
			r.FormValue("user_sexe")}
			fmt.Println(traitement)
		http.Redirect(w, r, "/user/display", http.StatusMovedPermanently)
	})

	http.HandleFunc("/user/display", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "display", traitement)
	})

	rootDoc, _ := os.Getwd()
	fileserver := http.FileServer(http.Dir(rootDoc + "/asset"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))

	http.ListenAndServe("localhost:8080", nil)

}

type promo struct {
	Nom               string
	Filière           string
	Niveau            int
	Nombre_d_étudiant int
	ListeEtudiant     []Etudiant
}

type Etudiant struct {
	NomE  string
	NomFE string
	Age   int
	Genre string
}

type Donnes struct {
	Nom    string
	Prenom string
	Anniv  string
	Sexe   string
}
