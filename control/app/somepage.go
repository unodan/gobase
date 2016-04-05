package app

import (
	"net/http"

	"github.com/unodan/cache"
)

func Somepage(w http.ResponseWriter, r *http.Request, uri string, ca *cache.Store) {
	var (
		args = map[string]string{}
	)
	Templates["somepage"].ExecuteTemplate(w, "somepage.html", &Page{
		ID:      ca.Get("ID").(string),
		User:    nil,
		Slug:    Slug(uri),
		Title:   "SomePage",
		Head:    &Head{"MyApp", "http", "localhost", "8080", uri},
		Header:  nil,
		Sidenav: nil,
		Content: nil,
		Sidebar: nil,
		Footer:  &Footer{"Footer"},
		Args:    &args},
	)
}
