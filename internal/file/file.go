package file

import (
	"io/ioutil"
	"path"
	"path/filepath"

	"github.com/pkg/errors"
)

// GetDir get all files from directory
func GetDir(dir string) ([]Dir, error) {
	ff := []Dir{}
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, errors.Wrap(err, "error reading dir")
	}

	for _, f := range files {
		ff = append(ff, Dir{
			Name: f.Name(),
			Path: path.Join(dir, f.Name()),
		})
	}

	return ff, nil
}

// GetContent from file
func GetContent(file string) (File, error) {
	bb, err := ioutil.ReadFile(file)
	if err != nil {
		return File{}, errors.Wrap(err, "error reading file")
	}

	f := File{
		Name:    filepath.Base(file),
		Content: string(bb),
	}

	return f, nil
}
