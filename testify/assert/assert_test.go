package assert

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

type Object  struct {
	Value string
}

func TestStopTest(t *testing.T) {
	assert :=assert.New(t)
	assert.Equal(123, 1123, "they should be equal")
}

func TestSomething(t *testing.T) {
	assert := assert.New(t)

	// assert equality
	assert.Equal(123, 123, "they should be equal")

	// assert inequality
	assert.NotEqual(123, 456, "they should not be equal")


	object := &Object{Value:"Something"}
	// assert for nil (good for errors)
	// assert.Nil(object)

	// assert for not nil (good when you expect something)
	if assert.NotNil(object) {
		// now we know that object isn't nil, we are safe to make
		// further assertions without causing any errors
		assert.Equal("Something", object.Value)
	}

}
