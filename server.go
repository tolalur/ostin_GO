package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type viewData struct{
 
    openIETag template.HTML
    closeIETag string
}

func main() {

	// отдаем статику
	fs := http.FileServer(http.Dir("static"))
	fmt.Println(fs)
	http.Handle("/", fs)

	http.HandleFunc("/dev", func(w http.ResponseWriter, r *http.Request) {
		htmlFiles := []string{
            "templates/index.html", 
            "templates/show_in_browser.html"}

        tmpl, _ := template.ParseFiles(htmlFiles...)
        data := viewData{
            openIETag: "<!--[if (gte mso 9)|(IE)]><table align=\"center\" cellspacing=\"0\" cellpadding=\"0\" border=\"0\" width=\"600\" style=\"width: 600px;\"><tr><td style=\"padding-top: 0px; padding-bottom: 0px; padding-right: 0px; padding-left: 0px; margin-top: 0px; margin-right: 0px; margin-bottom: 0px; margin-left: 0px;\"><![endif]-->",
            closeIETag: "<!--[if (gte mso 9)|(IE)]></td></tr></table><![endif]-->",
        }
        // data := viewData{
        //     openIETag: "Привет",
        //     closeIETag: "Привет2",
        // }
		tmpl.ExecuteTemplate(w, "index", data)
	})

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)
}
