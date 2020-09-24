package breadcrumbs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	assert := assert.New(t)

	bb := Create("testdata/foo/bar/bazz")

	assert.Equal(4, len(bb))

	assert.EqualValues(Breadcrumbs{
		Name: "testdata",
		Path: "testdata",
	}, bb[0])

	assert.EqualValues(Breadcrumbs{
		Name: "foo",
		Path: "testdata/foo",
	}, bb[1])

	assert.EqualValues(Breadcrumbs{
		Name: "bar",
		Path: "testdata/foo/bar",
	}, bb[2])

	assert.EqualValues(Breadcrumbs{
		Name: "bazz",
		Path: "testdata/foo/bar/bazz",
	}, bb[3])
}
