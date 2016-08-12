package graph

import (
	"io"
	"fmt"

	"github.com/wcharczuk/go-chart"
)

type LineService struct {

}

type Line struct {
	height int 
	width int
	values []*Value
}

func (p *Line) Height() int {
	return p.height
}

func (p *Line) Width() int {
	return p.width
}

func (p *Line) Values() []*Value {
	return p.values
}

func NewLineGraph(h int, w int, vsx []float64, vsy []float64, ls []string) *Line {	
	if shouldIsnsert0x0(vsx, vsy) {
		vsx = append(vsx, 0)
		vsy = append(vsy, 0)
	}
	vs := append(vsx, vsy...)
	numberOfValues := len(vs)
	values := make([]*Value, numberOfValues)
	j := 0
	for i := 0; i < numberOfValues; i++ {
		values[i] = &Value{
			Value: vs[i],
		}
		values[i].Label = fmt.Sprintf("%.2f", vs[i])
		if  i >= len(values)/2 && j < len(ls) {
			values[i].Label = ls[j]
			j = j +1
		}
	}	
	
	return &Line{
		height: h,
		width: w,
		values: values,
	}
}

func (s *LineService) Build(p *Line, w io.Writer) {
	values := p.Values()
	valuesXToChart := make([]float64, len(values))
	ticks := make([]chart.Tick, len(values)/2)
	j := 0
	for i := 0; i < len(values); i++ {
		valuesXToChart[i] = values[i].Value
		if i >= len(values)/2 {
			ticks[j] = chart.Tick{values[i].Value, values[i].Label}
			j = j+1
		} 
	}

	graph := chart.Chart{
		Width:  p.Width(),
		Height: p.Height(),
		XAxis: chart.XAxis{
			Style: chart.Style{
				Show: true,
			},
		},
		YAxis: chart.YAxis{
			Style: chart.Style{
				Show: true,
			},
			Ticks: ticks,
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: valuesXToChart[0:len(valuesXToChart)/2],
				YValues: valuesXToChart[len(valuesXToChart)/2:],
			},
		},
	}

	graph.Render(chart.PNG, w)
}

func shouldIsnsert0x0 (vsx []float64, vsy []float64) bool {
	for i := 0; i < len(vsx); i++ {
		if vsx[i] == 0 || vsy[i] == 0 {
			return false
		}
	}

	return true
}
