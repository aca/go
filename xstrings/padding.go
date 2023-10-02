package xstrings

import (
	"fmt"
	"strings"

	"golang.org/x/exp/constraints"
)

type Number interface {
    constraints.Integer | constraints.Float | constraints.Complex
}

func ZFill[T Number](v T, length int) string {
	return fmt.Sprintf("%0" + fmt.Sprintf("%d", length) + "v", v)
}

func RJustSpace[T Number | string](v T, length int) string {
	return RJust(v, length, ' ')
}

func LJustSpace[T Number | string](v T, length int) string {
	return LJust(v, length, ' ')
}

func RJust[T Number | string](v T, length int, padding byte) string {
	var str string
    switch p := any(v).(type) {
    case string:
		str = p
    default:
		str = fmt.Sprintf("%v", p)
    }

	paddingLen := length - len(str) 
	if paddingLen <= 0 {
		return str
	}
	return strings.Repeat(string(padding), paddingLen) + str
}

func LJust[T Number | string](v T, length int, padding byte) string {
	var str string
    switch p := any(v).(type) {
    case string:
		str = p
    default:
		str = fmt.Sprintf("%v", p)
    }

	paddingLen := length - len(str) 
	if paddingLen <= 0 {
		return str
	}

	return str + strings.Repeat(string(padding), paddingLen)
}
