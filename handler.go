package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"text/template"
	"time"

	blackfriday "github.com/russross/blackfriday/v2"
)

type HomeHandler struct {
	TemplatePath string
	TemplateName string
}

func (hh *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p, err := os.ReadFile(filepath.Join(hh.TemplatePath, hh.TemplateName))
	if err != nil {
		log.Printf("%v", err)
		w.WriteHeader(500)
		return
	}
	_, err = w.Write(p)
	if err != nil {
		log.Printf("%v", err)
		w.WriteHeader(500)
	}
}

type BlogHandler struct {
	TemplatePath string
	TemplateName string
	Config       string
}

func (bh *BlogHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	blog, err := NewTimeSortedBlog(bh.Config)
	if err != nil {
		fmt.Printf("%v", err)
		w.WriteHeader(500)
		return
	}

	for i, b := range blog.Posts {

		md, err := os.ReadFile("./" + b.File)
		if err != nil {
			fmt.Printf("%v", err)
			w.WriteHeader(500)
			return
		}

		html := blackfriday.Run([]byte(md))

		b.Content = string(html)
		blog.Posts[i].Content = string(html)
	}

	tf, err := os.ReadFile(filepath.Join(bh.TemplatePath, bh.TemplateName))
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
}

type BlogIndexHandler struct {
	TemplatePath string
	TemplateName string
	Config       string
}

func (bih *BlogIndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	blog, err := NewTimeSortedBlog(bih.Config)
	if err != nil {
		fmt.Printf("%v", err)
		w.WriteHeader(500)
		return
	}

	tf, err := os.ReadFile(filepath.Join(bih.TemplatePath, bih.TemplateName))
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

func NewTimeSortedBlog(filepath string) (Blog, error) {
	bcj, err := os.ReadFile(filepath)
	if err != nil {
		return Blog{}, err
	}

	var blog Blog
	err = json.Unmarshal(bcj, &blog)
	if err != nil {
		fmt.Printf("%v", err)
		return Blog{}, err
	}

	sort.Slice(blog.Posts, func(i, j int) bool {
		return blog.Posts[i].Date.After(blog.Posts[j].Date)
	})

	return blog, nil
}
