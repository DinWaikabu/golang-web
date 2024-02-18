package handlers

import (
	"app-web/entity"
	"html/template"
	"log"
	"net/http"
	"path"
)

// Homehandler digunakan untuk handle aktifitas yang akan dimuat untuk page home
func HomeHandler(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles(path.Join("views","index.html"), path.Join("views", "layout.html"))

	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
		return
	}

	// data:= map[string]interface{}{
	// 	"title": "Belajar Golang WEB",
	// 	"content": "Saya sedang belajar golang web",
	// }

	data := []entity.Product{
				{ID: 1, Nama:  "Laptop Asus", Price:  2500000, Stock:  10},
				{ID: 2, Nama:  "Laptop Hp", Price: 3000000, Stock:   8},
				{ID: 3, Nama:  "Macbook", Price:  4000000, Stock:   6},
				}


	err = tmpl.Execute(w, data)

	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
		return
	}
}


func AboutHandler(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles(path.Join("views", "about.html"), path.Join("views", "layout.html"))
	if err != nil {
		http.Error(w, "Error loading the about page.", http.StatusInternalServerError) 
		return 
	}
	
	data := struct {
		Title string
		Content string
	}{
		Title: "About Us",
		Content: "Kami adalah perusahaan yang bergerak	 dibidang jual beli barang-barang Content",
	}
	err = tmpl.Execute(w, data) // Execute Template dari layout dengan data diatas
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to load the about page.", http.StatusInternalServerError)
		return	 
	}
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	productId := vars.Get("id")
	if productId == "" {
		w.WriteHeader(400)
		w.Write([]byte("Missing Product Id"))
		return
	}
	data := map[string]interface{}{
		"content": productId,
	}

	tmpl, err := template.ParseFiles(path.Join("views","product.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Template not found", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		return
	}
}

func PostGet(w http.ResponseWriter, r *http.Request) {

	method := r.Method
	switch method {
	case "GET":
		w.Write([]byte("This is a GET request"))	
	case "POST":
		w.Write([]byte("This is a POST request"))
	default:
		w.WriteHeader(405) // Method Not Allowed
	}
}


// Handler form
func Form(w http.ResponseWriter, r *http.Request)  {
	if r.Method == "GET" {
		tmpl, err := template.ParseFiles(path.Join("views", "form.html"), path.Join("views", "layout.html"))
		if err != nil {
			log.Println(err)
			http.Error(w, "error is happening, keep calm", http.StatusInternalServerError)
			return
		}

		err=tmpl.Execute(w,nil)
		if err != nil {	
			log.Println(err)
			http.Error(w, "error is happening, keep calm", http.StatusInternalServerError)
			return
		}
		return
	}
	http.Error(w, "error is happening, keep calm", http.StatusBadRequest)
}

// Handler process

func Process(w  http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err:= r.ParseForm()
		if err!=nil{
			log.Println(err)
			http.Error(w,"can't parse the form ",http.StatusInternalServerError)
			return
		}
		name:=r.Form.Get("name")
		message := r.Form.Get("message")

		data := map[string]interface {}{
			"name": name,
			"message": message,
		}

		tmpl, err := template.ParseFiles(path.Join("views", "result.html"), path.Join("views", "layout.html"))
		if err != nil {
			log.Println(err)
			http.Error(w, "error is happening, keep calm", http.StatusInternalServerError)
			return
		}

		err=tmpl.Execute(w,data)
		if err != nil {	
			log.Println(err)
			http.Error(w, "error is happening, keep calm", http.StatusInternalServerError)
			return
		}
		return
	}
	http.Error(w, "Only POST method is supported.", http.StatusBadRequest)
}