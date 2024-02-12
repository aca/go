package filepathx

import "testing"

func TestAll(t *testing.T) {
	if FilenameWithComma("asd.tar.gz") != "asd.tar" {
		t.Fatal("FilenameWithComma failed")
	}

	if Filename("asd.tar.gz") != "asd" {
		t.Fatal("Filename failed")
	}

	if ExtWithComma("asd.tar.gz") != ".tar.gz" {
		t.Fatal("ExtWithComma failed")
	}

	if Ext("asd.tar.gz") != ".gz" {
		t.Fatal("Ext failed")
	}
}
