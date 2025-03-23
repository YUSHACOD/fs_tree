package dir

import (
	"fmt"
	"strings"

	"github.com/YUSHAOCOD/fs_tree/file"
)

type Directory struct {
	Path      string
	Name      string
	ChildDir  []Directory
	ChildFile []file.File
}

func PrintDirectory(dir *Directory, depth int) {
	indent := strings.Repeat("    ", depth)
	fmt.Println(indent + "path: " + dir.Path)
	fmt.Println(indent + "name: " + dir.Name)
	fmt.Println(indent + "child files: ")
	for _, f := range dir.ChildFile {
		file.PrintFile(&f, depth+1)
	}
	fmt.Println(indent + "child dirs: ")
	fmt.Println()
	for _, d := range dir.ChildDir {
		PrintDirectory(&d, depth+1)
	}
}

func CreateDirectory(path string, rem_path string, depth int) Directory {
	rem_path = strings.TrimLeft(rem_path, "/")
	nodes := strings.Split(rem_path, "/")

	path = strings.TrimLeft(path, "/")
	full_nodes := strings.Split(path, "/")
	dir_name := full_nodes[depth]
	dir_path := "/"
	for i := range depth + 1 {
		dir_path = dir_path + full_nodes[i] + "/"
	}

	new_dirs := []Directory{}
	files := []file.File{}

	if len(nodes) <= 1 {
		f := file.GetFile(path)
		// file.PrintFile(&f, )
		files = append(files, f)
	} else {
		new_path := "/"

		for i := 1; i < len(nodes)-1; i++ {
			// fmt.Println(nodes[i])
			new_path = new_path + nodes[i] + "/"
		}
		new_path = new_path + nodes[len(nodes)-1]

		directory := CreateDirectory(path, new_path, depth+1)
		// PrintDirectory(&directory, 0)
		new_dirs = append(new_dirs, directory)
	}

	return Directory{
		Path:      dir_path,
		Name:      dir_name,
		ChildDir:  new_dirs,
		ChildFile: files,
	}
}
