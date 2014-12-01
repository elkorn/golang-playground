package presentation

import (
	"io"
	"text/template"

	"github.com/elkorn/golang-playground/concurrency/schedule-calendar/domain"
)

type Page struct {
	Title string
	Body  string
	Users map[string]domain.User
}

func MkPage(title, body string, users map[string]domain.User) *Page {
	return &Page{
		Title: title,
		Body:  body,
		Users: users,
	}
}

func (self *Page) Render(filepath string, w io.Writer) {
	t, _ := template.ParseFiles(filepath)
	t.Execute(w, self)
}
