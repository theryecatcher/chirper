package web

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/theryecatcher/chirper/userd/userdpb"

	"github.com/julienschmidt/httprouter"
	"github.com/theryecatcher/chirper/contentd/contentdpb"
	"github.com/theryecatcher/chirper/web/session"
	"github.com/theryecatcher/chirper/web/views"
)

// TweetsGet displays the tweets in the timeline
func (ws *Web) TweetsGet(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Get session
	sess := session.Instance(r)

	UIDs := make([]string, 0)
	userID := fmt.Sprintf("%s", sess.Values["id"])

	user, usrErr := ws.userDaemon.GetUser(context.Background(), &userdpb.GetUserRequest{
		UID: userID,
	})
	if usrErr != nil {
		log.Println(usrErr)
	}

	UIDs = append(UIDs, userID)
	UIDs = append(UIDs, user.User.FollowingUID...)
	fmt.Println(UIDs)

	tweets, twtErr := ws.contentDaemon.GetTweetsByUser(context.Background(), &contentdpb.GetTweetsByUserRequest{
		UID: UIDs,
	})
	if twtErr != nil {
		log.Println(twtErr)
		tweets = tweets
	}

	allUsers, userderr := ws.userDaemon.GetAllFollowers(context.Background(), &userdpb.FollowerDetailsRequest{
		UID: userID,
	})
	if userderr != nil {
		log.Println(userderr)
	}

	// Need to Parallellize
	for _, tweet := range tweets.Tweets {
		usr, err := ws.userDaemon.GetUser(context.Background(), &userdpb.GetUserRequest{
			UID: tweet.PosterUID,
		})
		if err != nil {
			log.Println(err)
			tweet.PosterUID = "Anonymous"
		} else {
			tweet.PosterUID = usr.User.Name
		}
	}

	// Display the view
	v := view.New(r)
	v.Name = "tweets/read"
	v.Vars["name"] = sess.Values["name"]
	v.Vars["tweets"] = tweets.Tweets
	v.Vars["followers"] = allUsers.Followers
	v.Render(w)
}

// TweetsCreateGet displays the Tweet post page
func (ws *Web) TweetsCreateGet(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Get session
	// sess := session.Instance(r)

	// Display the view
	v := view.New(r)
	v.Name = "tweets/create"
	v.Render(w)
}

// TweetsCreatePost handles the Tweet post form submission
func (ws *Web) TweetsCreatePost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Get session
	sess := session.Instance(r)

	// Validate with required fields
	if validate, missingField := view.Validate(r, []string{"tweet"}); !validate {
		sess.AddFlash(view.Flash{"Field missing: " + missingField, view.FlashError})
		sess.Save(r, w)
		ws.TweetsCreateGet(w, r, nil)
		return
	}

	// Get form values
	content := r.FormValue("tweet")

	userID := fmt.Sprintf("%s", sess.Values["id"])

	// Do Database Insert
	_, err := ws.contentDaemon.NewTweet(context.Background(), &contentdpb.NewTweetRequest{
		Content:   content,
		PosterUID: userID,
	})
	// Error if there is a problem with the query
	if err != nil {
		log.Println(err)
		sess.AddFlash(view.Flash{"An error occurred on the server. Please try again later.", view.FlashError})
		sess.Save(r, w)
	} else {
		sess.AddFlash(view.Flash{"Tweet Posted!", view.FlashSuccess})
		sess.Save(r, w)
		http.Redirect(w, r, "/tweets", http.StatusFound)
		return
	}

	// Display the same page
	ws.TweetsCreateGet(w, r, nil)
}
