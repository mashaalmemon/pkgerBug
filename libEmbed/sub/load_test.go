package sub

import "testing"
import "github.com/stretchr/testify/assert"

// Check that loading of embedded file working
func Test_OpenEmbeddedFile(t *testing.T) {
	err := OpenEmbeddedFile()

	assert.NoError(t, err, "Expect no error opening embedded file")
}

// Check that loading of embedded file when using module prefix
func Test_OpenEmbeddedFileModulePrefix(t *testing.T) {
	err := OpenEmbeddedFileModulePrefix()

	assert.NoError(t, err, "Expect no error opening embedded file")
}
