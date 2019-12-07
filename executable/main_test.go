package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Check that loading embedded file from another module works
func Test_OpenModuleEmbeddedFile(t *testing.T) {
	err := OpenModuleEmbeddedFile()

	assert.NoError(t, err, "Expect no error opening embedded file")
}

// Check that loading embedded file from another module using its own module prefix works
func Test_OpenModuleEmbeddedFileModulePrefix(t *testing.T) {
	err := OpenModuleEmbeddedFileModulePrefix()

	assert.NoError(t, err, "Expect no error opening embedded file")
}
