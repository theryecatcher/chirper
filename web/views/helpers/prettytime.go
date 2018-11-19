package helpers

import (
	"html/template"
	"time"
)

// PrettyTime returns a template.FuncMap
// * PRETTYTIME outputs a nice time format
func PrettyTime() template.FuncMap {
	f := make(template.FuncMap)

	f["PRETTYTIME"] = func(t int64) string {
		time := time.Unix(t, 0)
		return time.Format("3:04 PM 01/02/2006")
	}

	return f
}
