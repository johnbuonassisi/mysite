package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"text/template"
	"time"

	blackfriday "gopkg.in/russross/blackfriday.v2"
)

func TestBlackFriday(t *testing.T) {

	md := "# Hello World!"
	html := blackfriday.Run([]byte(md))
	fmt.Printf("%v", string(html))

	tf, err := ioutil.ReadFile("./static/blog.html")
	if err != nil {
		t.Fatalf("%v", err)
	}

	temp, err := template.New("blog").Parse(string(tf))
	if err != nil {
		t.Fatalf("%v", err)
	}

	err = temp.Execute(os.Stdout, fmt.Sprintf("%s", string(html)))
	if err != nil {
		t.Fatalf("%v", err)
	}
}

type TestBlog struct {
	Posts []TestPost
}

type TestPost struct {
	Name string
	Date time.Time
}

func TestTemplateRange(t *testing.T) {

	data := TestBlog{Posts: []TestPost{{Name: "1", Date: time.Now()}, {Name: "2", Date: time.Now()}}}
	htmlTemplate := "{{range .Posts}} <h1>{{.Name}}<h1> \n{{end}}"

	template, err := template.New("t").Parse(htmlTemplate)
	if err != nil {
		t.Fatalf("%v", err)
	}

	err = template.Execute(os.Stdout, data)
	if err != nil {
		t.Fatalf("%v", err)
	}

}
