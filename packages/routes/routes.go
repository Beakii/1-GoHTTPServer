package routes

import (
	"log"
	"net/http"
	"text/template"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", homePage)
	mux.HandleFunc("/alt", altHome)

	return mux
}

func homePage(w http.ResponseWriter, r *http.Request) {
	//Get HTML file
	var filePath = "./templates/home.html"
	var fileName = "home.html"
	t, err := template.ParseFiles(filePath)
	if err != nil {
		log.Fatal(err)
		return
	}
	t.ExecuteTemplate(w, fileName, nil)
}
func altHome(w http.ResponseWriter, r *http.Request) {
	//Get HTML file
	var filePath = "./templates/altHome.html"
	var fileName = "altHome.html"
	t, err := template.ParseFiles(filePath)
	if err != nil {
		log.Fatal(err)
		return
	}
	t.ExecuteTemplate(w, fileName, nil)
}
