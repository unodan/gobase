package app

import (
	"net/http"

	"github.com/unodan/cache"
)

func Help(w http.ResponseWriter, r *http.Request, uri string, ca *cache.Store) {
	var (
		title    = "Help Page"
		template = "index"

		args = map[string]string{
			"Title":   "Help Page Title.",
			"Message": "This is the content area for the help page.",
		}

		page = &Page{
			ID:      ca.Get("ID").(string),
			User:    nil,
			Slug:    Slug(uri),
			Title:   title,
			Head:    &Head{title, "http", "localhost", "8080", uri},
			Header:  nil,
			Sidenav: &[]Sidenav{{"/", "Home Page"}, {"/somepage", "Some Page"}},
			Content: nil,
			Sidebar: nil,
			Footer:  &Footer{"Footer"},
			Args:    &args,
		}
	)

	Templates[template].ExecuteTemplate(w, template+".html", page)
}
