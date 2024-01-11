package views

import (
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
)

func (t Template) Execute(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := t.htmlTpl.Execute(w, data)
	if err != nil {
		log.Printf("executing template: %v", err)
		http.Error(w, "Error executing the template", http.StatusInternalServerError)
		return
	}
}

type Template struct {
	htmlTpl *template.Template
}

func ParseFS(fs fs.FS, pattern ...string) (Template, error) {
	htmlTpl, err := template.ParseFS(fs, pattern...)
	if err != nil {
		return Template{}, fmt.Errorf("parsing template: %w", err)
	}
	return Template{
		htmlTpl: htmlTpl,
	}, nil
}

//func Parse(filepath string) (Template, error) {
//	htmlTpl, err := template.ParseFiles(filepath)
//	if err != nil {
//		return Template{}, fmt.Errorf("parsing template: %w", err)
//	}
//	return Template{
//		htmlTpl: htmlTpl,
//	}, nil
//}

func Must(t Template, err error) Template {
	if err != nil {
		panic(err)
	}
	return t
}
