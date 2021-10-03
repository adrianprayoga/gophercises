package main

import (
	cyoa "cyoa/story"
	"flag"
	"html/template"
	"log"
	"net/http"
)

type handler struct {
	s      cyoa.Story
	t      *template.Template
	pathFn func(r *http.Request) string
}

type HandlerOption func(*handler)

func (h handler) storyHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		storyArc := h.pathFn(r)
		s, ok := h.s[storyArc]
		if !ok {
			log.Printf("Error reading storyArc %s", storyArc)
			http.NotFound(w, r)
			return
		}

		h.t.Execute(w, s)
	}
}

func WithTemplate(t *template.Template) HandlerOption {
	return func(h *handler) {
		h.t = t
	}
}

func WithPathFunc(pathFn func(r *http.Request) string) HandlerOption {
	return func(h *handler) {
		h.pathFn = pathFn
	}
}

func defaultParser(r *http.Request) string {
	path := r.URL.Path
	if path == "" || path == "/" {
		path = "/intro"
	}
	return path[1:]
}

func storyPathParser(r *http.Request) string {
	base := "/story/"
	path := r.URL.Path
	if path == base {
		path = base + "intro"
	}
	return path[len(base):]
}

func defaultTemplate() *template.Template {
	template, err := template.ParseFiles("htmls/template.html")
	if err != nil {
		log.Fatal("error reading templates")
	}
	return template
}

func createHandler(s cyoa.Story, opts ...HandlerOption) handler {
	h := handler{s, defaultTemplate(), defaultParser}
	for _, f := range opts {
		f(&h)
	}
	return h
}

func main() {
	storyPtr := flag.String("s", "storybook/story1.json", "storybook json")
	flag.Parse()

	// 1. Read JSON
	story := cyoa.GetJson(storyPtr)

	mux := defaultMux()
	h := createHandler(story, WithPathFunc(storyPathParser))

	mux.HandleFunc("/story/", h.storyHandler())
	http.ListenAndServe(":8080", mux)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	return mux
}
