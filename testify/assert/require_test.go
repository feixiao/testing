package assert

import (
	"testing"
	"github.com/stretchr/testify/require"
)


func TestRequire(t *testing.T) {
	require :=require.New(t)

	// assert equality
	require.Equal(123, 1123, "they should be equal")

	// assert inequality
	require.NotEqual(123, 456, "they should not be equal")

}
