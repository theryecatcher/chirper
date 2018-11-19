package web

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Static maps static files
func (ws *Web) Static(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	http.ServeFile(w, r, r.URL.Path[1:])
}
