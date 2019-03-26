package main

import (
	"fmt"
	"html/template"
	"net/http"
	"./config"
)


func main() {

	// отдаем статику
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	http.HandleFunc("/dev", func(w http.ResponseWriter, r *http.Request) {
		
		htmlFiles := config.Templates()
		data := config.GetData()

        tmpl, _ := template.ParseFiles(htmlFiles...)


		tmpl.ExecuteTemplate(w, "index", data)
	})

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)
}
