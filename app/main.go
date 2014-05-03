package hlilje

import (
    "html/template"
    "net/http"
)

type Index struct {
	Title string
}

// Panics if error is non-nil
var indexTemplate = template.Must(template.New("index").ParseFiles("tmpl/index.html"))

func init() {
    http.HandleFunc("/", root)
}

func root(w http.ResponseWriter, r *http.Request) {
	index := Index {
		Title: "hlilje",
	}
    err := indexTemplate.Execute(w, index)

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
