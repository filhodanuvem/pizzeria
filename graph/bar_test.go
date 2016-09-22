package graph

import (
	"os"
	"testing"
)

func TestDefaultBarColors(t *testing.T) {
	filename := "./bartest2.png"
	file, err := os.Create(filename)
	defer file.Close()
	defer os.Remove(filename)
	if err != nil {
		t.Fatalf(err.Error())
	}

	g := NewBarGraph(200, 200, []float64{80, 15, 5}, []string{"cash", "debit", "credit"})
	service := new(BarService)
	service.Build(g, file)

	assertImages("../_images/bar_test1.png", filename, t)
}

func TestBarColorsRedGreenBlue(t *testing.T) {
	filename := "./bartest2.png"
	file, err := os.Create(filename)
	defer file.Close()
	defer os.Remove(filename)
	if err != nil {
		t.Fatalf(err.Error())
	}

	g := NewBarGraph(200, 200, []float64{80, 15, 5}, []string{"cash", "debit", "credit"})
	g.Colors = []string{"f00", "0f0", "00f"}
	service := new(BarService)
	service.Build(g, file)

	assertImages("../_images/bar_test2.png", filename, t)
}

func TestBarBigFile(t *testing.T) {
	filename := "./bartest3.png"
	file, err := os.Create(filename)
	defer file.Close()
	defer os.Remove(filename)
	if err != nil {
		t.Fatalf(err.Error())
	}

	g := NewBarGraph(600, 800, []float64{90, 10}, []string{"pizza", "candy"})
	service := new(BarService)
	service.Build(g, file)

	assertImages("../_images/bar_test3.png", filename, t)
}
