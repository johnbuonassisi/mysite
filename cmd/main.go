package main

import (
	"log"
	"net/http"
)

const templatePath string = "../template"
const staticPath string = "../static"
const blogConfig string = "../blog-config.json"

func main() {

	http.Handle("/", &HomeHandler{TemplatePath: templatePath, TemplateName: "home.html"})

	http.Handle("/blog/", &BlogHandler{TemplatePath: templatePath,
		TemplateName: "blog.html",
		Config:       blogConfig})

	http.Handle("/blog/index/", &BlogIndexHandler{TemplatePath: templatePath,
		TemplateName: "blog-index.html",
		Config:       blogConfig})

	fs := http.FileServer(http.Dir(staticPath))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	log.Println("Listening on :80...")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}
