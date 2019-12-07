package sub

import (
	"fmt"
	"github.com/markbates/pkger"
	"io"
	"os"
)

type Empty struct{}

// open embedded file without module prefix
func OpenEmbeddedFile() error {
	return open("/sub/HelloWorld.txt")
}

// open embedded file using module prefix
func OpenEmbeddedFileModulePrefix() error {
	return open("pkgerBug/libEmbed:/sub/HelloWorld.txt")
}

// open embedded file
func open(filePath string) error {
	f, err := pkger.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		return err
	}

	fmt.Println("Opening Embedded File: ", filePath)
	fmt.Println("Name: ", info.Name())
	fmt.Println("Size: ", info.Size())
	fmt.Println("Mode: ", info.Mode())
	fmt.Println("ModTime: ", info.ModTime())

	if _, err := io.Copy(os.Stdout, f); err != nil {
		return err
	}
	fmt.Printf("\n")
	fmt.Println("")

	return nil
}
