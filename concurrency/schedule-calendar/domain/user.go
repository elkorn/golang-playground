package domain

import "html/template"

type User struct {
	Name     string
	Times    map[int]bool
	DateHTML template.HTML
}
