package dir_test

import (
	"fmt"
	"testing"

	"github.com/YUSHAOCOD/fs_tree/dir"
	"github.com/YUSHAOCOD/fs_tree/file"
)

func TestDirectoryCreation(t *testing.T) {
	// "/user/Documents/reports/2025/finance/summary.pdf"
	// "/user/Documents/reports/2025/hr/employee_list.xlsx"
	// "/user/Music/Albums/Rock/Classic/led_zeppelin_iv.flac"
	// "/user/Music/Playlists/workout.m3u"
	// "/user/Videos/Movies/SciFi/Interstellar_2014.mkv"
	// "/user/Videos/Tutorials/programming/python_basics.mp4"
	// "/user/Pictures/Vacation/2024/Italy/Rome/colosseum.jpg"
	// "/user/Pictures/Screenshots/screenshot_2025-03-23.png"
	// "/user/Downloads/software/linux/ubuntu.iso"
	// "/user/Downloads/books/programming/rust_book.epub"
	// "/user/Projects/WebApp/backend/src/database/models/user.py"
	// "/user/Projects/WebApp/frontend/assets/css/styles.css"
	// "/user/Projects/GameDev/Unity/Assets/Scripts/PlayerController.cs"
	// "/user/.config/fish/config.fish"
	// "/user/.local/share/nvim/session.vim"

	input := "/user/Documents/reports/2025/finance/summary.pdf"
	input2 := "/user/Music/Albums/Rock/Classic/led_zeppelin_iv.flac"
	directory := dir.CreateDirectory(input, input, 0)

	root := dir.Directory {
		Path: "/",
		Name: "/",
		ChildDir: []dir.Directory{directory},
		ChildFile: []file.File{},
	}

	fmt.Println("Simple")
	fmt.Println()
	dir.PrintDirectory(&root, 0)
	fmt.Println()

	root.ChildDir = append(root.ChildDir, dir.AddDir(input2, root.ChildDir[0]))

	fmt.Println()
	fmt.Println()
	dir.PrintDirectory(&root, 0)
	fmt.Println()
}
