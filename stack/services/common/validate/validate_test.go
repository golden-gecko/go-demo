package validate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_String(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(NonEmptyString(""), false)
    assert.Equal(NonEmptyString("abc"), true)
}
