package web

import (
	"net/http"

	"github.com/distsys-project/web/session"
	"github.com/distsys-project/web/views"
)

// func (ws *Web) Index(w http.ResponseWriter, req *http.Request) {
// 	w.Write([]byte("Under construction"))
// }
// Index displays the home page
func (ws *Web) Index(w http.ResponseWriter, r *http.Request) {
	// Get session
	session := session.Instance(r)

	if session.Values["id"] != nil {
		// Display the view
		v := view.New(r)
		v.Name = "index/auth"
		v.Vars["first_name"] = session.Values["first_name"]
		v.Render(w)
	} else {
		// Display the view
		v := view.New(r)
		v.Name = "index/anon"
		v.Render(w)
		return
	}
}
