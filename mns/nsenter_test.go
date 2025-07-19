package mns

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNSEnter(t *testing.T) {
	err := nsEnter()
	require.NoError(t, err)
}
