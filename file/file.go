package file

import (
	"fmt"
	"strings"
)

type File struct {
	Path string
	Name string
}

func GetFile(path string) File {
	name := path[strings.LastIndex(path, "/") + 1: ]
	return File{
		Path: path,
		Name: name,
	}
}

func PrintFile(f *File, depth int) {
	indent := strings.Repeat("    ", depth)
	fmt.Println(indent + "path: " + f.Path)
	fmt.Println(indent + "name: " + f.Name)
}
