package server

import (
	"net/http"
	"os"
	"path"
	"text/template"

	"github.com/apex/log"
	"github.com/wesleimp/sfs/internal/file"
)

var (
	tmplDir  = template.Must(template.ParseFiles("internal/tmpl/dir.html"))
	tmplFile = template.Must(template.ParseFiles("internal/tmpl/file.html"))
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.WithField("path", r.URL.Path).Info("received request")
	location := path.Join(".", r.URL.Path)

	if !isDir(location) {
		f, _ := file.GetContent(location)
		tmplFile.Execute(w, f)
		return
	}

	dirs, _ := file.GetDir(path.Join(".", r.URL.Path))
	tmplDir.Execute(w, dirs)
}

func isDir(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		log.WithError(err).Error("error reading file")
	}

	return info.IsDir()
}
