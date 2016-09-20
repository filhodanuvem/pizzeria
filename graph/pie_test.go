package graph

import (
	"testing"
	"os"
)

func TestDefaultPieColors(t *testing.T) {
    filename := "./pietest2.png"
    file, err := os.Create(filename)
    defer file.Close()
    defer os.Remove(filename)
    if err != nil {
        t.Fatalf(err.Error())
    }

    g := NewPieGraph(200, 200, []float64{80, 15, 5}, []string{"cash", "debit", "credit"})
    service := new(PieService)
    service.Build(g, file)

    assertImages("../_images/pie_test1.png", filename,  t)
}

func TestColorsRedGreenBlue(t *testing.T) {
    filename := "./pietest2.png"
    file, err := os.Create(filename)
    defer file.Close()
    defer os.Remove(filename)
    if err != nil {
        t.Fatalf(err.Error())
    }

    g := NewPieGraph(200, 200, []float64{80, 15, 5}, []string{"cash", "debit", "credit"})
    g.Colors = []string{"f00", "0f0", "00f"}
    service := new(PieService)
    service.Build(g, file)

    assertImages("../_images/pie_test2.png", filename,  t)
}

func TestBigFile(t *testing.T) {
    filename := "./pietest2.png"
    file, err := os.Create(filename)
    defer file.Close()
    defer os.Remove(filename)
    if err != nil {
        t.Fatalf(err.Error())
    }

    g := NewPieGraph(800, 600, []float64{90, 10}, []string{"pizza", "candy"})
    service := new(PieService)
    service.Build(g, file)

    assertImages("../_images/pie_test3.png", filename,  t)
}