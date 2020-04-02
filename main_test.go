package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"text/template"

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
