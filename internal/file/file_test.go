package file

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFiles(t *testing.T) {
	assert := assert.New(t)

	files, err := GetDir("./testdata")
	assert.NoError(err)

	assert.Equal(3, len(files))

	assert.Equal("content.golden", files[0].Name)
	assert.Equal("testdata/content.golden", files[0].Path)

	assert.Equal("file1.golden", files[1].Name)
	assert.Equal("testdata/file1.golden", files[1].Path)

	assert.Equal("file2.golden", files[2].Name)
	assert.Equal("testdata/file2.golden", files[2].Path)
}

func TestGetContent(t *testing.T) {
	assert := assert.New(t)

	file, err := GetContent("./testdata/content.golden")
	assert.NoError(err)
	assert.Equal("content.golden", file.Name)
	assert.Equal("Hey FOO", file.Content)
}
