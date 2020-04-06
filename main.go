package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
	"text/template"
	"time"

	blackfriday "gopkg.in/russross/blackfriday.v2"
)

func main() {

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	// create endpoint for main page
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

	http.HandleFunc("/blog/index/", func(w http.ResponseWriter, r *http.Request) {
		bcj, err := ioutil.ReadFile("./blog.json")
		if err != nil {
			fmt.Printf("%v", err)
			return
		}

		var blog Blog
		err = json.Unmarshal(bcj, &blog)
		if err != nil {
			fmt.Printf("%v", err)
			return
		}

		sort.Slice(blog.Posts, func(i, j int) bool {
			return blog.Posts[i].Date.After(blog.Posts[j].Date)
		})

		for i, b := range blog.Posts {

			md, err := ioutil.ReadFile("./" + b.File)
			if err != nil {
				fmt.Printf("%v", err)
				w.WriteHeader(500)
				return
			}

			html := blackfriday.Run([]byte(md))

			b.Content = string(html)
			blog.Posts[i].Content = string(html)
		}

		tf, err := ioutil.ReadFile("./blog-index.html")
		if err != nil {
			fmt.Printf("%v", err)
			w.WriteHeader(500)
			return
		}

		template, err := template.New("blog-index").Parse(string(tf))
		if err != nil {
			fmt.Printf("%v", err)
			w.WriteHeader(500)
			return
		}

		err = template.Execute(w, blog)
		if err != nil {
			fmt.Printf("%v", err)
			w.WriteHeader(500)
			return
		}
	})

	// create endpoint for blog page
	http.HandleFunc("/blog/", func(w http.ResponseWriter, r *http.Request) {

		bcj, err := ioutil.ReadFile("./blog.json")
		if err != nil {
			fmt.Printf("%v", err)
			return
		}

		var blog Blog
		err = json.Unmarshal(bcj, &blog)
		if err != nil {
			fmt.Printf("%v", err)
			return
		}

		sort.Slice(blog.Posts, func(i, j int) bool {
			return blog.Posts[i].Date.After(blog.Posts[j].Date)
		})

		for i, b := range blog.Posts {

			md, err := ioutil.ReadFile("./" + b.File)
			if err != nil {
				fmt.Printf("%v", err)
				w.WriteHeader(500)
				return
			}

			html := blackfriday.Run([]byte(md))

			b.Content = string(html)
			blog.Posts[i].Content = string(html)
		}

		tf, err := ioutil.ReadFile("./blog.html")
		if err != nil {
			fmt.Printf("%v", err)
			w.WriteHeader(500)
			return
		}

		template, err := template.New("blog").Parse(string(tf))
		if err != nil {
			fmt.Printf("%v", err)
			w.WriteHeader(500)
			return
		}

		err = template.Execute(w, blog)
		if err != nil {
			fmt.Printf("%v", err)
			w.WriteHeader(500)
			return
		}
	})

	log.Println("Listening on :80...")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}

type Blog struct {
	Posts []Post `json:"posts"`
}

type Post struct {
	File        string    `json:"file"`
	Reference   string    `json:"reference"`
	Type        string    `json:"type"`
	Title       string    `json:"title"`
	Date        time.Time `json:"date"`
	Author      string    `json:"author"`
	Description string    `json:"description"`
	Content     string
}

func (p *Post) FormattedDate() string {
	return p.Date.Format("January 2, 2006")
}
