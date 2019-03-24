package main

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
)

func main() {
	buf := new(bytes.Buffer)
	tmpl, err := template.ParseFiles("templates/index.html", "templates/head.html")

	if err != nil {
		fmt.Println("Ошибка парсинга")
		fmt.Printf("", err)
	}

	tmpl.ExecuteTemplate(buf, "index", "");
	templateStringx := buf.String();

	file, err := os.Create("./static/index.html")
     
    if err != nil{
        fmt.Println("Unable to create file:", err) 
        os.Exit(1) 
	}
	
    defer file.Close() 
    file.WriteString(templateStringx)
     
    fmt.Println("Done.")

}
