package decimalx

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDecimal_Add2(t *testing.T) {
	x := New(1, -7)
	one := New(1, -7)

	for i := 0; i < 10000000-1; i++ {
		x = x.Add(one)
	}
	require.Equal(t, x.String(), "1.0000000")

	x = x.Minus(one)
	require.Equal(t, x.String(), "0.9999999")
	require.Equal(t, one.String(), "0.0000001")
}

func TestDecimal_Add(t *testing.T) {
	x := New(1, -5)
	y := New(1, -5)

	require.Equal(t, x.Add(y).String(), "0.00002")
}

func TestDecimal_IsSmaller(t *testing.T) {
	x := New(2, -5)
	y := New(1, -5)
	require.True(t, x.IsBigger(y))
	require.False(t, x.IsSmaller(y))
}

func TestDecimal_Minus(t *testing.T) {
	x := New(1, -5)
	require.Equal(t, Zero.Minus(x).String(), "-0.00001")
}

func TestDecimal_Div(t *testing.T) {
	x := New(2, -5)
	y := New(2, 0)
	require.Equal(t, y.String(), "2")
	require.Equal(t, x.String(), "0.00002")
	require.Equal(t, x.Div(y).String(), "0.00001")
}
