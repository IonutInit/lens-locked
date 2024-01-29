package views

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/gorilla/csrf"
	"html/template"
	"io"
	"io/fs"
	"lens-locked/context"
	"lens-locked/models"
	"log"
	"net/http"
)

func (t Template) Execute(w http.ResponseWriter, r *http.Request, data interface{}, errs ...error) {
	tpl, err := t.htmlTpl.Clone()
	if err != nil {
		log.Printf("cloning template: %v", err)
		http.Error(w, "There was an error rendering the page", http.StatusInternalServerError)
		return
	}
	errMsgs := errMessages(errs...)
	tpl = tpl.Funcs(
		template.FuncMap{
			"csrfField": func() template.HTML {
				return csrf.TemplateField(r)
			},
			"currentUser": func() *models.User { return context.User(r.Context()) },
			//"errors": func() []string {
			//	var msgs []string
			//	for _, err := range errs {
			//		var pubErr public
			//		if errors.As(err, &pubErr) {
			//			msgs = append(msgs, pubErr.Public())
			//		} else {
			//			fmt.Println(err)
			//			msgs = append(msgs, "Something went wrong")
			//		}
			//	}
			//	return errMsgs
			//},
			"errors": func() []string {
				return errMsgs
			},
		},
	)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	var buf bytes.Buffer
	err = tpl.Execute(&buf, data)
	if err != nil {
		log.Printf("executing template: %v", err)
		http.Error(w, "There was an error executing the template", http.StatusInternalServerError)
		return
	}
	io.Copy(w, &buf)
}

type Template struct {
	htmlTpl *template.Template
}

func ParseFS(fs fs.FS, patterns ...string) (Template, error) {
	tpl := template.New(patterns[0])
	tpl = tpl.Funcs(
		template.FuncMap{
			"csrfField": func() (template.HTML, error) {
				return "", fmt.Errorf("scrfField not implemented")
			},
			"currentUser": func() (*models.User, error) { return nil, fmt.Errorf("currentUser not implemented") },
			"errors": func() []string {
				return nil
			},
		},
	)
	tpl, err := tpl.ParseFS(fs, patterns...)
	if err != nil {
		return Template{}, fmt.Errorf("parsing template: %w", err)
	}
	return Template{
		htmlTpl: tpl,
	}, nil
}

func Must(t Template, err error) Template {
	if err != nil {
		panic(err)
	}
	return t
}

type public interface {
	Public() string
}

func errMessages(errs ...error) []string {
	var msgs []string
	for _, err := range errs {
		var pubErr public
		if errors.As(err, &pubErr) {
			msgs = append(msgs, pubErr.Public())
		} else {
			fmt.Println(err)
			msgs = append(msgs, "Something went wrong")
		}
	}
	return msgs
}
