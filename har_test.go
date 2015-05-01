package har

import (
	"testing"
)

func TestLoadFile(t *testing.T) {
	_, err := LoadFile("testdata/test.har")
	if err != nil {
		t.Fatal(err)
	}
}
