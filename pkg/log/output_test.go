package log

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestOutputWriter(t *testing.T) {
	stdout := File("./l/file.test")
	file := File("./file.test")

	assert.IsType(t, os.Stdout, stdout)
	assert.IsType(t, &os.File{}, file)

	os.Remove("./file.test")
}
