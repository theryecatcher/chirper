package web

import (
	"bytes"
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/theryecatcher/chirper/userd/userdpb"
	"github.com/theryecatcher/chirper/web/session"
	"github.com/theryecatcher/chirper/web/views"
)

// RegisterGet displays the register page
func (ws *Web) RegisterGet(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Get session
	// sess := session.Instance(r)

	// Display the view
	v := view.New(r)
	v.Name = "register/register"
	// Refill any form fields
	view.Repopulate([]string{"first_name", "last_name", "email"}, r.Form, v.Vars)
	v.Render(w)
}

// RegisterPost handles the registration form submission
func (ws *Web) RegisterPost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Get session
	sess := session.Instance(r)

	// Validate with required fields
	if validate, missingField := view.Validate(r, []string{"first_name", "last_name", "email", "password"}); !validate {
		sess.AddFlash(view.Flash{"Field missing: " + missingField, view.FlashError})
		sess.Save(r, w)
		ws.RegisterGet(w, r, nil)
		return
	}

	// Get form values
	firstName := r.FormValue("first_name")
	lastName := r.FormValue("last_name")
	email := r.FormValue("email")
	password := r.FormValue("password")

	// Get database result
	var buffer bytes.Buffer
	buffer.WriteString(firstName)
	buffer.WriteString(" ")
	buffer.WriteString(lastName)

	_, err := ws.userDaemon.NewUser(context.Background(), &userdpb.NewUserRequest{
		Name:     buffer.String(),
		Email:    email,
		Password: password,
	})
	// Will only error if there is a problem with the query
	if err != nil {
		if err.Error() == "rpc error: code = Unknown desc = User already exists" {
			sess.AddFlash(view.Flash{"Account already exists for: " + email, view.FlashError})
			sess.Save(r, w)
		} else {
			ws.logger.Println(err)
			sess.AddFlash(view.Flash{"An error occurred on the server. Please try again later.", view.FlashError})
			sess.Save(r, w)
		}
	} else {
		sess.AddFlash(view.Flash{"Account created successfully for: " + email, view.FlashSuccess})
		sess.Save(r, w)
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	// Display the page
	ws.RegisterGet(w, r, nil)
}
