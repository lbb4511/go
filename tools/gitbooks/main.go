package main

// Copyright 2018 <lbb4511{AT}126.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ingore

//
// 批量浏览多个gitbook电子书
//
// Example:
//	gitbooks
//	gitbooks dir
//
// Help:
//	gitbooks -h
//

import (
	"bufio"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type Book struct {
	Name string
	Href string
	Url  string
	Img  string
	Key  int
}

var (
	books = make(map[string]*Book)
	tmpl  = `<html>
    <head>
        <title>Web</title>
    </head>
    <body>
        <p>Books:</p>
        {{range .}}
            <a href="{{.Href}}">
                <img src="{{.Img}}" alt="{{.Name}}" width="300" height="400"/>
            </li>
        {{end}}
    </body>
    </html>`
)

const usage = `
Usage: gitbooks dir [gitbook build]
gitbooks -h

Example:
  gitbooks
  gitbooks dir

  Report bugs to <lbb4511{AT}126.com>.
`

func main() {

	if len(os.Args) < 2 || os.Args[1] == "-h" {
		fmt.Fprintln(os.Stderr, usage[1:len(usage)-1])
		os.Exit(0)
	}
	dir := os.Args[1]

	var urlo string
	key := 0
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err == nil && info.IsDir() {
			if relpath, err := filepath.Rel(dir, path); err == nil {
				if strings.Contains(relpath, "node_modules") {
					return nil
				} else if strings.Contains(relpath, "_book") {
					var url string
					for _, rel := range strings.Split(relpath, "/") {
						if rel == "_book" {
							break
						}
						url += "/" + rel
					}
					if urlo != url {
						urlo = url
						key++
						books[url] = &Book{
							Key:  key,
							Url:  url,
							Name: getName(dir + url),
							Href: url + "/_book",
							Img:  url + "/_book/cover.jpg",
						}
					}
				}
			} else {
				return err
			}
		}
		return err
	})

	http.Handle("/", http.FileServer(http.Dir(dir)))

	http.HandleFunc("/book", handler)

	log.Print("Serving book on http://localhost:8080/book")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {

	tmpl, err := Template()

	if err != nil {
		fmt.Fprintf(w, "ParseFiles: %v", err)
		return
	}

	err = tmpl.Execute(w, books)

	if err != nil {
		fmt.Fprintf(w, "Execute: %v", err)
	}
}

func getName(url string) string {
	f, err := os.Open(url + "/README.md")
	defer f.Close()
	if nil == err {
		buff := bufio.NewReader(f)
		for {
			line, _, err := buff.ReadLine()
			if err != nil || io.EOF == err {
				break
			}
			return string(line[2:])
		}
	}
	return ""
}

func Template() (*template.Template, error) {
	return template.New("web").Funcs(template.FuncMap{
		"Str2html": func(str string) template.HTML {
			return template.HTML(str)
		},
		"Divide": func(num int) int {
			return num / 2
		},
		"Add": func(num int) int {
			return num + 100
		},
	}).Parse(tmpl)

}
