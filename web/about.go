package web

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/theryecatcher/chirper/web/views"
)

// AboutGet displays the About page
func (ws *Web) AboutGet(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Display the view
	v := view.New(r)
	v.Name = "about/about"
	v.Render(w)
}
