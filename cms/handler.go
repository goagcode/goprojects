package cms

import (
	"net/http"
	"strings"
	"time"
)

func ServePage(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimLeft(r.URL.Path, "/page/")
	if path == "" {
		http.NotFound(w, r)
		return
	}
	p := &Page{
		Title:   strings.ToTitle(path),
		Content: "Here is my page",
	}
	Tmpl.ExecuteTemplate(w, "page", p)
}

func ServePost(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimLeft(r.URL.Path, "/post/")
	if path == "" {
		http.NotFound(w, r)
		return
	}
	p := &Post{
		Title:   strings.ToTitle(path),
		Content: "Here is may page",
		Comments: []*Comment{
			&Comment{
				Author:        "Miguel Angel",
				Comment:       "Looks great!",
				DatePublished: time.Now(),
			},
		},
	}
	Tmpl.ExecuteTemplate(w, "post", p)
}

func HandleNew(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		Tmpl.ExecuteTemplate(w, "new", nil)

	case "POST":
		title := r.FormValue("title")
		content := r.FormValue("content")
		contentType := r.FormValue("content-type")
		r.ParseForm()

		if contentType == "page" {
			Tmpl.ExecuteTemplate(w, "page", &Page{
				Title:   title,
				Content: content,
			})
			return
		}

		if contentType == "post" {
			Tmpl.ExecuteTemplate(w, "post", &Post{
				Title:   title,
				Content: content,
			})
			return
		}
	default:
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
	}
}

func ServeIndex(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title:   "Go Project CMS",
		Content: "Welcome to our home pages",
		Posts: []*Post{
			&Post{
				Title:         "Hello world",
				Content:       "Go is awesome",
				DatePublished: time.Now(),
				Comments: []*Comment{
					&Comment{
						Author:        "Miguel Angel",
						Comment:       "I will come back",
						DatePublished: time.Now().Add(-time.Hour / 2),
					},
				},
			},
		},
	}
	Tmpl.ExecuteTemplate(w, "page", p)
}
