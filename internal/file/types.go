package file

type (
	// Dir contains some info from file
	Dir struct {
		Name string
		Path string
	}

	// File contains a file content
	File struct {
		Name    string
		Content string
	}
)
