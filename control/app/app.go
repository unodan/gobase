/*
# File app.go
# Author: Dan Huckson
# Date: 20160401
*/
package app

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/unodan/cache"
)

var (
	cStore    *cache.Store
	validPath *regexp.Regexp
	Templates map[string]*template.Template
)

type User struct {
	User_handle         string
	User_password       string
	User_first_name     string
	User_last_name      string
	User_email          string
	User_url            string
	User_registered     time.Time
	User_activation_key string
	User_status         string
}

type Page struct {
	ID      string
	User    *User
	Title   string
	Slug    string
	Head    *Head
	Header  *Header
	Sidenav *[]Sidenav
	Content *Content
	Sidebar *Sidebar
	Footer  *Footer
	Args    *map[string]string
	Errors  []string
}
type Head struct {
	Title  string
	Scheme string
	Host   string
	Port   string
	URI    string
}
type Header struct {
	Content string
}
type Sidenav struct {
	Url  string
	Text string
}
type Content struct {
	Content string
}
type Sidebar struct {
	Content string
}
type Footer struct {
	Content string
}

type Cache struct {
	ID       string
	Title    string
	Path     string
	HostIP   string
	HostName string
	HostPort string
	Scheme   string
}

func Bail(err int) int {
	return err
}
func Slug(uri string) string {
	i, j := strings.LastIndex(uri, "/"), strings.LastIndex(uri, path.Ext(uri))
	return uri[i+1 : j]
}

func Setup(c *Cache) *cache.Store {
	var (
		err     error
		path, _ = os.Getwd()
		ca      = cStore.New()
	)
	cacheData := Cache{
		ID:       "",
		Title:    "",
		Path:     path,
		HostIP:   "127.0.0.1",
		HostName: "localhost",
		HostPort: "8080",
		Scheme:   "http",
	}

	dflt := reflect.Indirect(reflect.ValueOf(&cacheData))
	if c != nil {
		vCache := reflect.Indirect(reflect.ValueOf(c))

		for i := 0; i < dflt.NumField(); i++ {
			switch fmt.Sprintf("%T", vCache.Field(i).Interface()) {
			case "string":
				if vCache.Field(i).Interface().(string) == "" {
					ca.Set(dflt.Type().Field(i).Name, dflt.Field(i).Interface().(string))
				} else {
					ca.Set(vCache.Type().Field(i).Name, vCache.Field(i).Interface().(string))
				}
				break
			}
		}
	}

	if err = LoadTemplates(); err != nil {
		log.Println("app, Error: application [ " + ca.Get("Title").(string) + " ] has failed to initialize, program aborted")
		log.Println("app,", err)
		os.Exit(Bail(2))
	}
	return ca
}
func MakeHandler(fn func(http.ResponseWriter, *http.Request, string, *cache.Store), uri string, cs *cache.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		validPath := regexp.MustCompile("^(" + uri + ")")
		if r.URL.Path == "/" {
			fn(w, r, "/signup", cs)
		} else if m := validPath.FindStringSubmatch(r.URL.Path); m != nil {
			fn(w, r, m[1], cs)
		} else {
			fn(w, r, "/help", cs)
		}
	}
}
func LoadTemplates() error {
	Templates = make(map[string]*template.Template)
	pages, err := filepath.Glob("templates/*.html")
	if err == nil {
		sections1, err := filepath.Glob("templates/forms/*.html")
		if err == nil {
			sections2, err := filepath.Glob("templates/sections/*.html")
			if err == nil {
				sections := append(sections1, sections2...)
				for _, p := range pages {
					files := append(sections, p)
					page := strings.Split(filepath.Base(p), ".")[0]
					Templates[page] = template.Must(template.ParseFiles(files...))
				}
				log.Println("app, Info: all templates have been loaded")
			} else {
				log.Println("app, Error: could not load templates (\"templates/sections/*.html\")")
			}
		} else {
			log.Println("app, Error: could not load templates (\"templates/forms/*.html\")")
		}
	} else {
		log.Println("app, Error: could not load templates (\"templates/*.html\")")
	}
	return err
}
