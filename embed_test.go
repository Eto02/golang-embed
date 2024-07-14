package main_test

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"log"
	"os"
	"testing"
)

//go:embed  version.txt
var version string

func Test_Name(t *testing.T) {
	fmt.Println(version)
}

//go:embed github.png
var image []byte

func TestByte(t *testing.T) {
	err := os.WriteFile("sample.png", image, fs.ModePerm)
	if err != nil {
		log.Print(err)
	}
}

//go:embed files/a.txt
//go:embed files/b.txt
var files embed.FS

func TestMultipleFiles(t *testing.T) {
	a, _ := files.ReadFile("files/a.txt")
	fmt.Println(string(a))

	b, _ := files.ReadFile("files/b.txt")
	fmt.Println(string(b))

	c, _ := files.ReadFile("files/c.txt")
	fmt.Println(string(c))
}

//go:embed files/*.txt
var path embed.FS

func TestPathMatcher(t *testing.T) {
	dirEntries, _ := path.ReadDir("files")
	for _, entry := range dirEntries {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			file, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println(string(file))
		}
	}
}
