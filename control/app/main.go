package app

import (
	"net/http"

	"github.com/unodan/cache"
)

func Main(w http.ResponseWriter, r *http.Request, uri string, ap *cache.Store) {
	var (
		args = map[string]string{}
	)
	Templates["main"].ExecuteTemplate(w, "main.html", &Page{
		ID:      ap.Get("ID").(string),
		User:    nil,
		Slug:    Slug(uri),
		Title:   "MainPage",
		Head:    &Head{"MyApp", "http", "localhost", "8080", uri},
		Header:  nil,
		Sidenav: nil,
		Content: nil,
		Sidebar: nil,
		Footer:  &Footer{"Footer"},
		Args:    &args},
	)
}
