package web

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/theryecatcher/chirper/web/session"
	"github.com/theryecatcher/chirper/web/views"
)

// Index displays the home page
func (ws *Web) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Get session
	session := session.Instance(r)
	ws.logger.Println(session.Values["id"])

	if session.Values["id"] != nil {
		// Display the view
		v := view.New(r)
		v.Name = "index/auth"
		v.Vars["name"] = session.Values["name"]
		v.Render(w)
	} else {
		// Display the view
		v := view.New(r)
		v.Name = "index/unauth"
		v.Render(w)
		return
	}
}
