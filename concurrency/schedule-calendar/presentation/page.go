package presentation

import "github.com/elkorn/golang-playground/concurrency/schedule-calendar/domain"

type Page struct {
	Title string
	Body  string
	Users map[string]domain.User
}
