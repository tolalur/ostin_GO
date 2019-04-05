package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
	"./config"
)


func main() {

	// отдаем статику
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	http.HandleFunc("/dev", func(w http.ResponseWriter, r *http.Request) {
		
		t0 := time.Now()
		
		htmlFiles := config.Templates()
		data := config.GetData()

        tmpl, _ := template.ParseFiles(htmlFiles...)


		tmpl.ExecuteTemplate(w, "index", data)
		
		fmt.Printf("Elapsed time: %v", time.Since(t0))

	})

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)
}
