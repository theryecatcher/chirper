package web

import (
	"context"
	"fmt"
	"net/http"

	"github.com/theryecatcher/chirper/userd/userdpb"

	"github.com/julienschmidt/httprouter"
	"github.com/theryecatcher/chirper/web/session"
	"github.com/theryecatcher/chirper/web/views"
)

// LoginGet displays the login page
func (ws *Web) LoginGet(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Get session
	// sess := session.Instance(r)

	// Display the view
	v := view.New(r)
	v.Name = "login/login"
	// Refill any form fields
	view.Repopulate([]string{"email"}, r.Form, v.Vars)
	v.Render(w)
}

// LoginPost handles the login form submission
func (ws *Web) LoginPost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Get session
	sess := session.Instance(r)

	// Validate with required fields
	if validate, missingField := view.Validate(r, []string{"email", "password"}); !validate {
		sess.AddFlash(view.Flash{"Field missing: " + missingField, view.FlashError})
		sess.Save(r, w)
		ws.LoginGet(w, r, nil)
		return
	}

	// Form values
	email := r.FormValue("email")
	password := r.FormValue("password")

	// Get database result
	result, err := ws.userDaemon.ValidateUser(context.Background(), &userdpb.ValidateUserRequest{
		Email:    email,
		Password: password,
	})

	ws.logger.Println(result)
	// Determine if user exists
	if err != nil {
		if err.Error() == "rpc error: code = Unknown desc = User not found" { // no user exists with that email
			ws.logger.Println(err)
			sess.AddFlash(view.Flash{"User is not registered", view.FlashWarning})
			sess.Save(r, w)
		} else if err.Error() == "rpc error: code = Unknown desc = Incorrect password" {
			sess.AddFlash(view.Flash{"Password is incorrect.", view.FlashWarning})
			sess.Save(r, w)
		} else {
			// Display error message
			ws.logger.Println(err)
			sess.AddFlash(view.Flash{"There was an error on the server. Please try again later.", view.FlashError})
			sess.Save(r, w)
		}
	} else {
		// Login successfully
		session.Empty(sess)
		sess.AddFlash(view.Flash{"Login successful!", view.FlashSuccess})
		sess.Values["id"] = result.User.UID
		sess.Values["email"] = email
		sess.Values["name"] = result.User.Name
		err := sess.Save(r, w)
		if err != nil {
			fmt.Println(err)
		}
		ws.logger.Println(sess.Values["id"])
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	// Show the login page again
	ws.LoginGet(w, r, nil)
}

// LogoutGet clears the session and logs the user out
func (ws *Web) LogoutGet(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Get session
	sess := session.Instance(r)

	// If user is authenticated empty session
	if sess.Values["id"] != nil {
		session.Empty(sess)
		sess.AddFlash(view.Flash{"Goodbye!", view.FlashNotice})
		sess.Save(r, w)
	}

	http.Redirect(w, r, "/", http.StatusFound)
}
