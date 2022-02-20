package web

import (
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"
)

func Init() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("./cmd/web/views/index.html")
		if err != nil {
			log.Println("Unexpected template err:", err.Error())
		}

		tmpl.Execute(w, nil)
	})

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = ":8080"
	} else if !strings.HasPrefix(":", port) {
		port = ":" + port
	}
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./cmd/web/views/public/"))))
	log.Println("Serving on " + port)
	http.ListenAndServe(port, nil)
}
