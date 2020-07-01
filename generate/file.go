package generate

import (
	"github.com/go-zookeeper/jute/parser"
)

// File is a parsed file
type File struct {
	Path string
	Doc  *parser.Doc
}

// ParseFile will parse a single file.
func ParseFile(filename string) (*File, error) {
	doc, err := parser.ParseFile(filename)
	if err != nil {
		return nil, err
	}

	return &File{
		Path: filename,
		Doc:  doc,
	}, nil
}

// ParseFiles will parse multiple files ready for generation.
func ParseFiles(filenames ...string) ([]*File, error) {
	files := []*File{}

	for _, file := range filenames {
		file, err := ParseFile(file)
		if err != nil {
			return nil, err
		}
		files = append(files, file)
	}
	return files, nil
}
