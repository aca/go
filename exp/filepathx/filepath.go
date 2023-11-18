package filepathx

import "path/filepath"

func Ext(path string) string {
	return filepath.Ext(path)
}
