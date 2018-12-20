package web

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/theryecatcher/chirper/contentd/contentdpb"
	"github.com/theryecatcher/chirper/userd/userdpb"
	"google.golang.org/grpc"

	"github.com/julienschmidt/httprouter"
)

type Web struct {
	srv *http.Server

	contentDaemon contentdpb.ContentdClient
	userDaemon    userdpb.UserdClient

	logger *log.Logger
}

func New(cfg *Config) (*Web, error) {

	loclLogger := log.New(os.Stderr, "[webServer] ", log.LstdFlags)

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	cntdConn, err := grpc.Dial("localhost:5445", opts...)
	if err != nil {
		loclLogger.Fatalf("failure while dialing: %v", err)
	}
	// defer cntdConn.Close()
	// Need to figure out adding this I keep getting error
	// rpc error: code = Canceled desc = grpc: the client connection is closing

	usrdConn, err := grpc.Dial("localhost:5446", opts...)
	if err != nil {
		loclLogger.Fatalf("failure while dialing: %v", err)
	}
	// defer usrdConn.Close()

	// mx := http.NewServeMux()
	r := httprouter.New()
	s := &http.Server{
		Addr:    cfg.HTTPAddr,
		Handler: r,
	}

	ws := &Web{
		srv:           s,
		contentDaemon: contentdpb.NewContentdClient(cntdConn),
		userDaemon:    userdpb.NewUserdClient(usrdConn),
		logger:        loclLogger,
	}

	r.GET("/static/*filepath", ws.Static)
	r.GET("/", ws.Index)
	r.GET("/about", ws.AboutGet)
	// Login/out
	r.GET("/login", ws.LoginGet)
	r.POST("/login", ws.LoginPost)
	r.GET("/logout", ws.LogoutGet)
	// Register
	r.GET("/register", ws.RegisterGet)
	r.POST("/register", ws.RegisterPost)
	// Tweets
	r.GET("/tweets", ws.TweetsGet)
	r.GET("/tweets/create", ws.TweetsCreateGet)
	r.POST("/tweets/create", ws.TweetsCreatePost)
	// Follow/Unfollow
	r.GET("/follow/follow/:uid", ws.FollowGet)
	r.GET("/follow/unfollow/:uid", ws.UnFollowGet)
	// r.GET("/follow/unfollow/:id", ws.FollowGet)

	return ws, nil
}

func (w *Web) Start() error {
	return w.srv.ListenAndServe()
}

func (w *Web) Shutdown(ctx context.Context) error {
	return w.srv.Shutdown(ctx)
}
