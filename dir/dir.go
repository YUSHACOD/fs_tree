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
		files = append(files, f)
	} else {
		new_path := "/"

		for i := 1; i < len(nodes)-1; i++ {
			new_path = new_path + nodes[i] + "/"
		}
		new_path = new_path + nodes[len(nodes)-1]

		directory := CreateDirectory(path, new_path, depth+1)
		new_dirs = append(new_dirs, directory)
	}

	return Directory{
		Path:      dir_path,
		Name:      dir_name,
		ChildDir:  new_dirs,
		ChildFile: files,
	}
}

func IsDirPresent(dirs []Directory, name string) int {
	for i, d := range dirs {
		if d.Name == name {
			return i
		}
	}
	return -1 
}

func AddDir(path string, directory Directory) Directory {
	new_dir := CreateDirectory(path, path, 0)

	new_path := strings.TrimLeft(path, "/")
	s := 0
	for i := 0; new_path[i] !=  '/'; i++ {
		s = i
	}
	new_path = new_path[s + 1:]
	

	if new_dir.Name == directory.Name {
		for _, child_dir := range new_dir.ChildDir {
			idx := IsDirPresent(directory.ChildDir, child_dir.Name) 
			if idx != -1 {
				directory.ChildDir[idx] = AddDir(new_path, directory.ChildDir[idx])
			} else {
				directory.ChildDir = append(directory.ChildDir, child_dir)
			}
		}

		return directory
	} else {
		return new_dir
	}
}
