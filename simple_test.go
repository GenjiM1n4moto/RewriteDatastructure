package rewritedatastructure

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTest1(t *testing.T) {
	test1 := test1()
	require.Equal(t, 3, test1)
}
