package fsutil

import (
	"path/filepath"
	"strings"
)

func Filename(filename string) string {
	return strings.TrimSuffix(filename, filepath.Ext(filename))
}

func Ext(filename string) string {
	return filepath.Ext(filename)
}

func ExtWithComma(filename string) string {
	filename = filepath.Base(filename)
	i := strings.Index(filename, ".")
	if i == -1 {
		return ""
	}
	return filename[i:]
}

func FilenameWithComma(filename string) string {
	filename = filepath.Base(filename)
	splitted := strings.Split(filename, ".")
	if len(splitted) == 0 {
		return ""
	} else {
		return splitted[0]
	}
}