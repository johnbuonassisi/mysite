package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"text/template"
	"time"

	blackfriday "gopkg.in/russross/blackfriday.v2"
)

func main() {

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		p, err := ioutil.ReadFile("./index.html")
		if err != nil {
			os.Exit(1)
		}
		_, err = w.Write(p)
		if err != nil {
			w.WriteHeader(500)
		}
	})

	// create an endpoint for each blog ...?
	// serve the latest blog for /blog

	http.HandleFunc("/blog", func(w http.ResponseWriter, r *http.Request) {

		/*
			var bc []BlogConfig
			err = json.Unmarshal(p, &bc)
			if err != nil {
				fmt.Printf("%v", err)
				w.WriteHeader(500)
				return
			}
		*/

		md, err := ioutil.ReadFile("./static/blog/hello.md")
		if err != nil {
			fmt.Printf("%v", err)
			w.WriteHeader(500)
			return
		}

		html := blackfriday.Run([]byte(md))
		fmt.Printf("%v", string(html))

		tf, err := ioutil.ReadFile("./blog.html")
		if err != nil {
			fmt.Printf("%v", err)
			w.WriteHeader(500)
			return
		}

		temp, err := template.New("blog").Parse(string(tf))
		if err != nil {
			fmt.Printf("%v", err)
			w.WriteHeader(500)
			return
		}

		err = temp.Execute(w, fmt.Sprintf("%s", string(html)))
		if err != nil {
			fmt.Printf("%v", err)
			w.WriteHeader(500)
			return
		}
	})

	// create a handler that will serve the
	// get list of blog posts
	// get the blog template

	log.Println("Listening on :80...")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}

type BlogConfig struct {
	Blog []Blog `json:"blog"`
}

type Blog struct {
	File  string    `json:"file"`
	Type  string    `json:"type"`
	Title string    `json:"title"`
	Date  time.Time `json:"date"`
}
