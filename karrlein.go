package main

import (
	"html/template"
	"log"
	"net/http"
)

type PageVariables struct {
	Name string
}

func main() {
	http.HandleFunc("/", HomePage)
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	HomePageVars := PageVariables{
		Name: "André Karrlein",
	}

	t, err := template.ParseFiles("template.html")
	if err != nil {
		log.Print("template parsing error: ", err)
	}
	err = t.Execute(w, HomePageVars)
	if err != nil {
		log.Print("template executing error: ", err)
	}
}
