/*
# File init.go
# Author: Dan Huckson
# Date: 20160401
*/
package gobase

import (
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"github.com/unodan/cache"
	"github.com/unodan/gobase/control/app"
)

func init() {
	var (
		appID   = "mwa"
		appName = "My Web Application"
	)
	log.Println("\n\n---  " + appName + "  " + strings.Repeat("-", 80-len(appName)-7) + "\n")
	log.Println("init, Info: application [ " + appName + " ] has started")

	r := mux.NewRouter()
	ca := app.Setup(&app.Cache{
		ID:    appID,
		Title: appName,
	})

	func(handlers map[string]func(http.ResponseWriter, *http.Request, string, *cache.Store)) {
		r.NotFoundHandler = app.MakeHandler(app.Help, "/help", ca)
		for route := range handlers {
			r.HandleFunc(route, app.MakeHandler(handlers[route], route, ca))
		}
		http.Handle("/", r)
		log.Println("init, Info: application [ " + ca.Get("Title").(string) + " ] has initialized successfully")
		log.Println("init, Info: server [ " + ca.Get("HostName").(string) + " ], waiting for requests on port [ " + ca.Get("HostPort").(string) + " ]")
	}(map[string]func(http.ResponseWriter, *http.Request, string, *cache.Store){
		"/":         app.Main,
		"/somepage": app.Somepage,
	})
}
