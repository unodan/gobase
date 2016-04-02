package app

import (
	"net/http"

	"github.com/unodan/cache"
)

func Help(w http.ResponseWriter, r *http.Request, uri string, ap *cache.Store) {
	var (
		args = map[string]string{
			"Title":   "The requested URL cannot be found.",
			"Message": "Do You Need Help?",
		}
	)
	Templates["help"].ExecuteTemplate(w, "help.html", &Page{
		ID:      ap.Get("ID").(string),
		User:    nil,
		Slug:    Slug(uri),
		Title:   "HelpPage",
		Head:    &Head{"Help", "http", "localhost", "8080", uri},
		Header:  nil,
		Sidenav: nil,
		Content: nil,
		Sidebar: nil,
		Footer:  &Footer{"Footer"},
		Args:    &args},
	)
}
