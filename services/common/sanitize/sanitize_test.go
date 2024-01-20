package sanitize

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_String(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(String(""), "")
    assert.Equal(String(" "), "")
	assert.Equal(String(" abc"), "abc")
	assert.Equal(String("abc "), "abc")
    assert.Equal(String(" abc "), "abc")
}