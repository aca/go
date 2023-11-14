package xstrings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZFill(t *testing.T) {
	{
		s := ZFill(3, 5)
		assert.Equal(t, "00003", s)
	}
	{
		s := ZFill(40.3, 5)
		assert.Equal(t, "040.3", s)
	}
}

func TestRJust(t *testing.T) {
	{
		s := RJustSpace(3, 5)
		assert.Equal(t, "    3", s)
		t.Log(s)
	}

	{
		s := RJustSpace("he", 5)
		assert.Equal(t, "   he", s)
		t.Log(s)
	}
}
