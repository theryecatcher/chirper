package web

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/theryecatcher/chirper/userd/userdpb"
	"github.com/theryecatcher/chirper/web/session"
	"github.com/theryecatcher/chirper/web/views"
)

// FollowGet allows user to follow other users
func (ws *Web) FollowGet(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Get session
	sess := session.Instance(r)

	userID := fmt.Sprintf("%s", sess.Values["id"])

	// params = gcontext.Get(r, "params").(httprouter.Params)
	followingID := ps.ByName("uid")

	// Get database result
	_, err := ws.userDaemon.FollowUser(context.Background(), &userdpb.FollowUserRequest{
		UID:          userID,
		FollowingUID: followingID,
	})
	// Will only error if there is a problem with the query
	if err != nil {
		log.Println(err)
		sess.AddFlash(view.Flash{"An error occurred on the server. Please try again later.", view.FlashError})
		sess.Save(r, w)
	} else {
		sess.AddFlash(view.Flash{"User Followed!", view.FlashSuccess})
		sess.Save(r, w)
	}

	http.Redirect(w, r, "/tweets", http.StatusFound)
	return
}

// UnFollowGet allows user to follow other users
func (ws *Web) UnFollowGet(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Get session
	sess := session.Instance(r)

	userID := fmt.Sprintf("%s", sess.Values["id"])

	// params = gcontext.Get(r, "params").(httprouter.Params)
	followerID := ps.ByName("uid")

	// Get database result
	_, err := ws.userDaemon.UnFollowUser(context.Background(), &userdpb.UnFollowUserRequest{
		UID:         userID,
		FollowedUID: followerID,
	})
	// Will only error if there is a problem with the query
	if err != nil {
		log.Println(err)
		sess.AddFlash(view.Flash{"An error occurred on the server. Please try again later.", view.FlashError})
		sess.Save(r, w)
	} else {
		sess.AddFlash(view.Flash{"User UnFollowed!", view.FlashSuccess})
		sess.Save(r, w)
	}

	http.Redirect(w, r, "/tweets", http.StatusFound)
	return
}
