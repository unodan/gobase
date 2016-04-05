package app

import (
	"net/http"

	"github.com/unodan/cache"
)

func Somepage(w http.ResponseWriter, r *http.Request, uri string, ca *cache.Store) {
	var (
		title    = "Some Page"
		template = "index"

		args = map[string]string{
			"Title":   "Some Page Title",
			"Message": "This is the content area for some page.",
		}

		page = &Page{
			ID:      ca.Get("ID").(string),
			User:    nil,
			Slug:    Slug(uri),
			Title:   title,
			Head:    &Head{title, "http", "localhost", "8080", uri},
			Header:  nil,
			Sidenav: &[]Sidenav{{"/", "Home Page"}, {"/help", "Help Page"}},
			Content: nil,
			Sidebar: nil,
			Footer:  &Footer{"Footer"},
			Args:    &args,
		}
	)

	Templates[template].ExecuteTemplate(w, template+".html", page)
}
