package graph

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func assertImages(expectedImage string, foundImage string, t *testing.T) {
	contentExpected, err := ioutil.ReadFile(expectedImage)
	if err != nil {
		t.Fatalf(err.Error())
	}

	content, err := ioutil.ReadFile(foundImage)
	if err != nil {
		t.Fatalf(err.Error())
	}

	if !bytes.Equal(content, contentExpected) {
		t.Fatalf("Image are different")
	}
}
