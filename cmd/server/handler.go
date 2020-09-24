package server

import (
	"net/http"
	"os"
	"path"
	"text/template"

	"github.com/apex/log"
	"github.com/wesleimp/sfs/internal/breadcrumbs"
	"github.com/wesleimp/sfs/internal/file"
)

var (
	tmplDir  = template.Must(template.ParseFiles("internal/tmpl/dir.html"))
	tmplFile = template.Must(template.ParseFiles("internal/tmpl/file.html"))
)

// DataDir holds all information about dirs
type DataDir struct {
	Breadcrumbs []breadcrumbs.Breadcrumbs
	Items       []file.Dir
}

// DataFile holds all information about dirs
type DataFile struct {
	Breadcrumbs []breadcrumbs.Breadcrumbs
	Name        string
	Content     string
}

func handler(w http.ResponseWriter, r *http.Request) {
	location := path.Join(".", r.URL.Path)
	log.WithField("path", location).Info("received request")

	bc := breadcrumbs.Create(location)

	if !isDir(location) {
		f, _ := file.GetContent(location)
		tmplFile.Execute(w, DataFile{
			Breadcrumbs: bc,
			Name:        f.Name,
			Content:     f.Content,
		})
		return
	}

	dirs, _ := file.GetDir(path.Join(".", r.URL.Path))
	tmplDir.Execute(w, DataDir{
		Breadcrumbs: bc,
		Items:       dirs,
	})
}

func isDir(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		log.WithError(err).Error("error reading file")
	}

	return info.IsDir()
}
