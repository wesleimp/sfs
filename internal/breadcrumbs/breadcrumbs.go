package breadcrumbs

import (
	"path"
	"strings"

	"github.com/apex/log"
)

// Create breadcrumbs
func Create(location string) []Breadcrumbs {
	paths := strings.Split(location, "/")
	log.WithField("paths", paths).Info("mounting breadcrumbs")
	bc := []Breadcrumbs{}

	if paths[0] != "." {
		bc = append(bc, Breadcrumbs{
			Name: ".",
			Path: "",
		})
	}

	for i, p := range paths {
		bc = append(bc, Breadcrumbs{
			Name: p,
			Path: path.Join(paths[:i+1]...),
		})
	}

	return bc
}
