package handlers

import (
	"html/template"
	"log"
	"net/http"
	"path"
)

// Homehandler digunakan untuk handle aktifitas yang akan dimuat untuk page home
func HomeHandler(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles(path.Join("views","index.html"))

	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
		return
	}

	data:= map[string]interface{}{
		"title": "Belajar Golang WEB",
		"content": "Saya sedang belajar golang web",
	}

	err = tmpl.Execute(w, data)

	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
		return
	}
}


func AboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("mario from Tanggerang"))
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	productId := vars.Get("id")
	if productId == "" {
		w.WriteHeader(400)
		w.Write([]byte("Missing Product Id"))
		return
	}
	w.Write([]byte("Product: " + productId))
}
