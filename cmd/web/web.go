package main

import (
	"encoding/json"

	"github.com/theryecatcher/chirper/web"
	"github.com/theryecatcher/chirper/web/session"
	"github.com/theryecatcher/chirper/web/views"
	"github.com/theryecatcher/chirper/web/views/helpers"
)

// config the settings variable
var config = &configuration{}

// configuration contains the application settings
type configuration struct {
	Session  session.Session `json:"Session"`
	Template view.Template   `json:"Template"`
	View     view.View       `json:"View"`
}

// ParseJSON unmarshals bytes to structs
func (c *configuration) ParseJSON(b []byte) error {
	return json.Unmarshal(b, &c)
}

func main() {

	JSONLoad("config.json", config)

	// Configure the session cookie store
	session.Configure(config.Session)

	// Setup the views
	view.Configure(config.View)
	view.LoadTemplates(config.Template.Root, config.Template.Children)
	view.LoadPlugins(
		helpers.TagHelper(config.View),
		helpers.PrettyTime())

	cfg := &web.Config{
		HTTPAddr: "localhost:80",
	}

	webSrv, err := web.New(cfg)
	if err != nil {
		panic(err)
	}

	err = webSrv.Start()
	if err != nil {
		panic(err)
	}
}
