package main

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"./config"
)

func main() {
	buf := new(bytes.Buffer)

	htmlFiles := config.Templates()
	data := config.GetData()
			
	tmpl, err := template.ParseFiles(htmlFiles...)

	if err != nil {
		fmt.Println("Ошибка парсинга")
		fmt.Printf("", err)
	}
		
	tmpl.ExecuteTemplate(buf, "index", data);
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
