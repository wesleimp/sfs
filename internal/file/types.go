package file

type (
	// Dir contains some info from file
	Dir struct {
		Name  string
		IsDir bool
		Path  string
	}

	// File contains a file content
	File struct {
		Name    string
		Content string
		Ext     string
	}
)
