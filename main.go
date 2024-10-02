package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/yuin/goldmark"
)

type Article struct {
	Title   string
	Content template.HTML
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/articles", articleHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	articles, err := getArticleList()
	if err != nil {
		http.Error(w, "Unable to load articles", http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/home.html"))
	err = tmpl.Execute(w, articles)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getArticleList() ([]string, error) {
	var articles []string
	files, err := os.ReadDir("articles")
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".md" {
			articles = append(articles, strings.TrimSuffix(file.Name(), ".md"))
		}
	}
	return articles, nil
}

func articleHandler(w http.ResponseWriter, r *http.Request) {
	articleName := strings.TrimPrefix(r.URL.Path, "/articles/")
	mdPath := filepath.Join("articles", articleName+".md")
	content, err := os.ReadFile(mdPath)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	var buf bytes.Buffer
	if err := goldmark.Convert(content, &buf); err != nil {
		http.Error(w, "Error processing article", http.StatusInternalServerError)
		return
	}

	article := Article{
		Title:   strings.ReplaceAll(articleName, "-", " "),
		Content: template.HTML(buf.String()),
	}

	tmpl := template.Must(template.ParseFiles("templates/article.html"))
	err = tmpl.Execute(w, article)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
