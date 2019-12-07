package main

import "pkgerBug/libEmbed/sub"

func main() {
	OpenModuleEmbeddedFile()
	OpenModuleEmbeddedFileModulePrefix()
}

// open file embedded in another module
func OpenModuleEmbeddedFile() error {
	return sub.OpenEmbeddedFile()
}

// open file embedded in another module and loading with module prefix
func OpenModuleEmbeddedFileModulePrefix() error {
	return sub.OpenEmbeddedFileModulePrefix()
}
