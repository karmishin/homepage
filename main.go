package main

import (
	"embed"
	"flag"
	"html/template"
	"log"
	"net/http"
)

var (
	//go:embed static/*
	staticFS embed.FS

	//go:embed templates/*
	templatesFS embed.FS
)

var templates = template.Must(template.ParseFS(templatesFS, "templates/index.html"))

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	templates.ExecuteTemplate(w, "index.html", nil)
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/search/" {
		http.NotFound(w, r)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	query := template.HTMLEscapeString(r.PostFormValue("query"))
	if query == "" {
		http.Error(w, "query is empty", http.StatusBadRequest)
		return
	}

	engine := r.PostFormValue("engine")
	if engine == "" {
		http.Error(w, "engine is empty", http.StatusBadRequest)
		return
	}

	searchUrl, err := getSearchUrl(query, engine)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Add("content-type", "text/plain")
	http.Redirect(w, r, searchUrl, http.StatusFound)
}

func main() {
	addr := flag.String("addr", ":8080", "Network address to listen on")
	flag.Parse()

	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.FS(staticFS))))
	http.HandleFunc("/", handler)
	http.HandleFunc("/search/", searchHandler)

	log.Println("Now listening on " + *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
