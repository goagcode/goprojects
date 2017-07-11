package main

import (
	"os"

	"github.com/miguellgt/goprojects/cms"
)

func main() {
	p := &cms.Page{
		Title:   "Hello, World",
		Content: "This is the body of our webpage",
	}

	cms.Tmpl.ExecuteTemplate(os.Stdout, "index", p)
}
