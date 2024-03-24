package main

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
)

type Page struct {
	Title string
	Body  []byte
}

var templates = template.Must(template.ParseFiles("./templates/top.html", "./templates/bottom.html", "./templates/edit.html", "./templates/view.html", "./templates/create.html", "./templates/index.html"))

var validPath = regexp.MustCompile("^/(edit|create|save|view|delete)/([a-zA-Z0-9]+)$")

var matches, _ = filepath.Glob("./txt/*.txt")

//go:embed static/css
var static embed.FS

func main() {
	http.HandleFunc("GET /", indexHandler)
	http.HandleFunc("GET /create/", createHandler)
	http.HandleFunc("POST /create/", makeHandler(saveHandler))
	http.HandleFunc("GET /view/", makeHandler(viewHandler))
	http.HandleFunc("GET /edit/", makeHandler(editHandler))
	http.HandleFunc("POST /save/", makeHandler(saveHandler))
	http.HandleFunc("DELETE /delete/", makeHandler(deleteHandler))
	http.Handle("GET /static/", http.FileServer(http.FS(static)))

	log.Fatal(http.ListenAndServe(":8080", nil))
	fmt.Println("Server started on port :8080")
}

func (p *Page) save() error {
	filename := "./txt/" + p.Title + ".txt"

	matches = append(matches, "./txt/"+p.Title+".txt")

	return os.WriteFile(filename, p.Body, 0600)
}

func (p *Page) delete() error {
	filename := "./txt/" + p.Title + ".txt"

	err := os.Remove(filename)
	fmt.Println(err)

	matches, _ = filepath.Glob("./txt/*.txt")

	return err
}

func loadPage(title string) (*Page, error) {
	filename := "./txt/" + title + ".txt"
	body, err := os.ReadFile(filename)

	return &Page{Title: title, Body: body}, err
}

func loadFile(title string) (*Page, error) {
	p, err := loadPage(title)

	return p, err
}

func renderPage(w http.ResponseWriter, tmpl string) {
	err := templates.ExecuteTemplate(w, tmpl+".html", matches)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
