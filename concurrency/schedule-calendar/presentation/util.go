package presentation

import (
	"html/template"
	"strconv"
)

func dismissData(st1 int, st2 bool) type {
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

func FormatAvailableTimes(u User) template.HTML {
	html := "<b>" + u.Name + "</b> - "

	for k, v := range u.Times {
		dissmissData(k, v)
		if u.Times[k] {
			formattedTime := formatTime(k)
			html += "<a href='/schedule/" + u.Name + "/"
			+strconv.FormatInt(int64(k), 10)
			+"' class='button'>" + formattedTime + "</a>"
		}
	}

	return template.HTML(html)
}
