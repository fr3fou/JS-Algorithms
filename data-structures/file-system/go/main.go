package main

import (
	"fmt"
	"github.com/fr3fou/go-fs/filesystem"
)

func main() {
	fs := filesystem.New()

	fs.CreateDir("usr")
	fs.CreateDir("usr/share")
	fs.CreateDir("usr/local")
	fs.ChangeDir("usr/local")

	fs.ChangeDir("../../../")

	items, _ := fs.ListDirectoryContents("usr")

	for k, v := range items {
		fmt.Printf("%v - %+v\n", k, v)
	}

	// fs.CreateDir("usr")
	// fs.ChangeDir("usr")

	// fmt.Println("fs: changed dir to: " + fs.PrintWorkingDirectory())

	// fs.CreateDir("share")
	// fs.ChangeDir("share")

	// fmt.Println("fs: changed dir to: " + fs.PrintWorkingDirectory())

	// fs.ChangeDir("../..")

	// fmt.Println("fs: changed dir to: " + fs.PrintWorkingDirectory())

	// fs.ChangeDir("usr/share/")

	// fmt.Println("fs: changed dir to: " + fs.PrintWorkingDirectory())

	// fs.ChangeDir("/usr/share")

	// fmt.Println("fs: changed dir to: " + fs.PrintWorkingDirectory())

	// fs.ChangeDir("..")

	// fmt.Println("fs: changed dir to: " + fs.PrintWorkingDirectory())

	// fs.ChangeDir("/")

	// fmt.Println("fs: changed dir to: " + fs.PrintWorkingDirectory())

	// fs.CreateFile("kernel", []byte("hello"))

	// content, _ := fs.ReadFile("kernel")
	// fmt.Println(string(content))

	// fs.ChangeDir("usr")

	// fmt.Println("fs: changed dir to: " + fs.PrintWorkingDirectory())

	// fs.CreateFile("nested", []byte("some content"))
	// fs.ChangeDir("../")

	// fmt.Println("fs: changed dir to: " + fs.PrintWorkingDirectory())

	// content, _ = fs.ReadFile("usr/nested")
	// fmt.Println(string(content))

	// fs.CreateFile("/usr/share/testing", []byte("testing root files"))

	// content, _ = fs.ReadFile("/usr/share/testing")
	// fmt.Println(string(content))

	// fs.CreateDir("/usr/share/fog")
	// fs.ChangeDir("/usr/share/fog")

	// fmt.Println("fs: changed dir to: " + fs.PrintWorkingDirectory())

	// // shouldn't work as the directories preceding it don't exist
	// fs.CreateDir("/usr/share/fog/this/shouldnt-work")
	// fs.ChangeDir("/usr/share/fog/this/shouldnt-work")

	// fs.DeleteFile("/usr/share/testing")
	// // shouldn't work as the file has been deleted
	// content, _ = fs.ReadFile("/usr/share/testing")

	// fs.ChangeDir("..")
	// fmt.Println("fs: changed dir to: " + fs.PrintWorkingDirectory())

	// items, _ := fs.ListDirectoryContents(fs.PrintWorkingDirectory())

	// for k, v := range items {
	// 	fmt.Printf("%v - %+v\n", k, v)
	// }

	// fs.DeleteDirectory("/usr/share/fog")
	// // shouldn't work as the directory has been deleted
	// items, err := fs.ListDirectoryContents("/usr/share/fog")
	// fmt.Println(err)

	// fs.CreateFile("kernel2", []byte("hey!"))
	// fs.EditFile("kernel2", []byte("hey again!"))
	// content, _ = fs.ReadFile("kernel2")
	// fmt.Println(string(content))

	// fmt.Println(fs.PrintWorkingDirectory())
	// items, _ = fs.ListDirectoryContents(fs.PrintWorkingDirectory())

	// for k, v := range items {
	// 	fmt.Printf("%v - %+v\n", k, v)
	// }
}
