package graph

import (
	"io"

	"github.com/wcharczuk/go-chart"
	"github.com/wcharczuk/go-chart/drawing"
)

type PieService struct {
}

type Pie struct {
	height int
	width  int
	Colors []string
	values []*Value
}

func (p *Pie) Height() int {
	return p.height
}

func (p *Pie) Width() int {
	return p.width
}

func (p *Pie) Values() []*Value {
	return p.values
}

func NewPieGraph(h int, w int, vs []float64, ls []string) *Pie {
	numberOfValues := len(vs)
	values := make([]*Value, numberOfValues)
	for i := 0; i < numberOfValues; i++ {
		values[i] = &Value{
			Value: vs[i],
			Label: ls[i],
		}
	}

	return &Pie{
		height: h,
		width:  w,
		values: values,
	}
}

func (s *PieService) Build(p *Pie, w io.Writer) {
	values := p.Values()
	valuesToChart := make([]chart.Value, len(values))
	for i := 0; i < len(values); i++ {
		valuesToChart[i] = chart.Value{
			Value: values[i].Value,
			Label: values[i].Label,
		}

		if len(p.Colors) > 0 {
			valuesToChart[i].Style = chart.Style{
				FillColor: drawing.ColorFromHex(p.Colors[i]),
			}
		}
	}

	graph := chart.PieChart{
		Width:  p.Width(),
		Height: p.Height(),
		Values: valuesToChart,
	}

	graph.Render(chart.PNG, w)
}
