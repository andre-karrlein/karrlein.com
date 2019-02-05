package main

import (
	"crypto/tls"
	"html/template"
	"log"
	"net/http"

	"golang.org/x/crypto/acme/autocert"
)

type PageVariables struct {
	Name string
}

func main() {
	certManager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("*.karrlein.com"), //Your domain here
		Cache:      autocert.DirCache("certs"),               //Folder for storing certificates
	}

	http.HandleFunc("/", MainPage)
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	server := &http.Server{
		Addr: ":443",
		TLSConfig: &tls.Config{
			GetCertificate: certManager.GetCertificate,
		},
	}

	go http.ListenAndServe(":80", certManager.HTTPHandler(nil))

	log.Fatal(server.ListenAndServeTLS("", "")) //Key and cert are coming from Let's Encrypt
}

func MainPage(w http.ResponseWriter, r *http.Request) {
	PageVars := PageVariables{
		Name: "André Karrlein",
	}

	t, err := template.ParseFiles("template.html")
	if err != nil {
		log.Print("template parsing error: ", err)
	}
	err = t.Execute(w, PageVars)
	if err != nil {
		log.Print("template executing error: ", err)
	}
}
