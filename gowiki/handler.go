package main

import (
	"fmt"
	"net/http"
	"strings"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	for m := 0; m < len(matches); m++ {
		matches[m] = strings.TrimSuffix(strings.TrimPrefix(matches[m], "txt/"), ".txt")
	}

	renderPage(w, "index")

}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path == "/save/" {
			fn(w, r, "")
		}

		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadFile(title)

	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
	}

	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, _ := loadFile(title)

	renderTemplate(w, "edit", p)
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	renderPage(w, "create")
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	if r.FormValue("title") != title {
		title = r.FormValue("title")
	}

	body := r.FormValue("body")

	p := &Page{Title: r.FormValue("title"), Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func deleteHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadFile(title)

	fmt.Println(p, err)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusNoContent)

	p.delete()
}
