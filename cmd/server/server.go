package server

import (
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"text/template"
	"time"

	"github.com/apex/log"
	"github.com/gorilla/mux"
	"github.com/urfave/cli/v2"
)

// Start server
func Start(c *cli.Context) error {
	port := c.String("port")

	tmplDir := template.Must(template.ParseFiles("internal/tmpl/dir.html"))
	tmplFile := template.Must(template.ParseFiles("internal/tmpl/file.html"))
	r := mux.NewRouter()

	r.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithField("path", r.URL.Path).WithField("raw path", r.URL.RawPath).Info("received request")
		location := path.Join(".", r.URL.Path)

		if !isDir(location) {
			bb, _ := ioutil.ReadFile(location)
			tmplFile.Execute(w, File{
				PageTitle: location,
				Code:      string(bb),
			})
			return
		}

		tmplDir.Execute(w, Dir{
			PageTitle: "Files",
			Files:     getFiles(path.Join(".", r.URL.Path)),
		})
	})

	srv := &http.Server{
		Handler:      r,
		Addr:         ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return srv.ListenAndServe()
}

type Dir struct {
	PageTitle string
	Files     []Files
}

type Files struct {
	Name string
	Path string
}

type File struct {
	PageTitle string
	Code      string
}

func getFiles(p string) []Files {
	ff := []Files{}
	files, _ := ioutil.ReadDir(p)
	for _, f := range files {
		ff = append(ff, Files{
			Name: f.Name(),
			Path: path.Join(p, f.Name()),
		})
	}
	return ff
}

func isDir(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		log.WithError(err).Error("error reading file")
	}

	return info.IsDir()
}
