//go:build re2_cgo

package docconv

import (
	"os"
	"testing"
)

func TestPDFHasImage_CannotExecuteCode(t *testing.T) {
	// Try to inject code by passing a bad file path.
	// If the code was successful it will create a file called foo in the working directory
	badFilePath := "$(id >> foo).pdf"
	got, err := PDFHasImage(badFilePath)
	if err != nil {
		t.Fatal(err)
	}
	if want := false; got != want {
		t.Errorf("got %v, want %v", got, want)
	}

	if got, want := fileExists("foo"), false; got != want {
		t.Errorf("got bad file exists, want not file to exist")
	}
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
