package presentation

import (
	"fmt"
	"html/template"
	"strconv"

	"github.com/elkorn/golang-playground/concurrency/schedule-calendar/domain"
)

func dismissData(st1 int, st2 bool) {
}

func formatTime(hour int) string {
	hourText := hour
	// ampm := "am"
	// if hour > 11 {
	// 	ampm = "am"
	// }
	//
	// if hour > 12 {
	// 	hourText = hour - 12
	// }

	return strconv.FormatInt(int64(hourText), 10) // + ampm
}

func FormatAvailableTimes(u domain.User) template.HTML {
	html := "<b>" + u.Name + "</b> - "

	for k, v := range u.Times {
		dismissData(k, v)
		if u.Times[k] {
			formattedTime := formatTime(k)
			html += fmt.Sprintf(
				"<a href='/schedule/%v/%d' class='button'>%s</a>",
				u.Name,
				strconv.FormatInt(int64(k), 10),
				formattedTime)
		}
	}

	return template.HTML(html)
}
