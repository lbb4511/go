package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type Package struct {
	Name     string
	NumFuncs int
	NumVars  int
}

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Getwd: %v", err)
	}
	log.Print("Work directory:", wd)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// tmpl := template.Must(template.ParseFiles(filepath.Join(wd, "main_v2.tmpl")))
		tmpl, err := template.ParseFiles(filepath.Join(wd, "index.tmpl"))
		if err != nil {
			fmt.Fprintf(w, "ParseFiles: %v", err)
			return
		}
		err = tmpl.Execute(w, &Package{
			Name:     "go-web",
			NumFuncs: 12,
			NumVars:  1200,
		})
		if err != nil {
			fmt.Fprintf(w, "Execute: %v", err)
			return
		}
	})
	http.HandleFunc("/v2", func(w http.ResponseWriter, r *http.Request) {
		pkg := &Package{
			Name:     "go-web",
			NumFuncs: 12,
			NumVars:  1200,
		}

		tmpl, err := template.New("indexF.tmpl").Funcs(template.FuncMap{
			"NumFuncs": func() int {
				return pkg.NumFuncs
			},
			"Str2html": func(str string) template.HTML {
				return template.HTML(str)
			},
			"Divide": func(num int) int {
				return num / 2
			},
			"Add": func(num int) int {
				return num + 100
			},
		}).ParseFiles(filepath.Join(wd, "indexF.tmpl"))
		if err != nil {
			fmt.Fprintf(w, "ParseFiles: %v", err)
			return
		}

		err = tmpl.Execute(w, map[string]interface{}{
			"Name":        pkg.Name,
			"NumFuncs":    pkg.NumFuncs,
			"NumVars":     pkg.NumVars,
			"NumVarsHTML": `<li>Number of functions: 1200</li>`,
			"Maps": map[string]map[string]string{
				"Level1": map[string]string{
					"Name": "go-web",
				},
			},
			"Nums": []int{1, 2, 3, 4, 5, 6, 7},
		})
		if err != nil {
			fmt.Fprintf(w, "Execute: %v", err)
			return
		}
	})

	log.Print("Starting server...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
