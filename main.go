package main

import (
	"ascii-art-web/art"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type ViewData struct {
	Message string
}

func main() {
	var data ViewData

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/output/result.txt" {
			tmpl1, err := template.ParseFiles("./output/result.txt")

			if err != nil {
				http.Error(w, "400 - Bad Request", http.StatusBadRequest)
				return
			}

			tmpl1.Execute(w, data)
			return
		}

		if r.URL.Path != "/" {
			http.Error(w, "404 not found.", http.StatusNotFound)
			return
		}

		if r.Method == "GET" {
			data = ViewData{
				Message: art.PrintAscii("Hello World", "standard"),
			}
		}

		if r.Method == "POST" {
			if err := r.ParseForm(); err != nil {
				fmt.Fprintf(w, "ParseForm() err: %v", err)
				return
			}
			stringToPrint := r.FormValue("string")
			texturesToUse := r.FormValue("Textures")

			data = ViewData{Message: art.PrintAscii(stringToPrint, texturesToUse)}
		}

		tmpl, err := template.ParseFiles("./templates/index.html")
		tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "400 - Bad Request", http.StatusBadRequest)
			return
		}
	})

	fmt.Println("Server is listening...")

	if http.ListenAndServe(":8000", nil) != nil {
		log.Fatalf("%v - Internal Server Error", http.StatusInternalServerError)
	}
}
