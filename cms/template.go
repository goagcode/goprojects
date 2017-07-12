package cms

import (
	"html/template"
	"os"
	"time"
)

var tmplPath = os.Getenv("GOPATH") + "/src/goprojects/cms"

// Tmpl is an exported variable
var Tmpl = template.Must(template.ParseGlob(tmplPath))

type Page struct {
	Title   string
	Content string
	Posts   []*Post
}

type Post struct {
	Title         string
	Content       string
	DatePublished time.Time
	Comments      []*Comment
}

type Comment struct {
	Author        string
	Comment       string
	DatePublished time.Time
}
