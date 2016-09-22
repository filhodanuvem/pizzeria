package graph

import (
	"io"

	"github.com/wcharczuk/go-chart"
	"github.com/wcharczuk/go-chart/drawing"
)

type BarService struct {
}

type Bar struct {
	height int
	width  int
	values []*Value
	Colors []string
}

func (b *Bar) Height() int {
	return b.height
}

func (b *Bar) Width() int {
	return b.width
}

func (b *Bar) Values() []*Value {
	return b.values
}

func NewBarGraph(h int, w int, vs []float64, ls []string) *Bar {
	numberOfValues := len(vs)
	values := make([]*Value, numberOfValues)
	for i := 0; i < numberOfValues; i++ {
		values[i] = &Value{
			Value: vs[i],
			Label: ls[i],
		}
	}

	return &Bar{
		height: h,
		width:  w,
		values: values,
	}
}

func (s *BarService) Build(b *Bar, w io.Writer) {
	values := b.Values()
	valuesToChart := make([]chart.Value, len(values))
	for i := 0; i < len(values); i++ {
		valuesToChart[i] = chart.Value{
			Value: values[i].Value,
			Label: values[i].Label,
		}

		if len(b.Colors) > 0 {
			valuesToChart[i].Style = chart.Style{
				FillColor:   drawing.ColorFromHex(b.Colors[i]),
				StrokeColor: drawing.ColorFromHex(b.Colors[i]),
			}
		}
	}

	graph := chart.BarChart{
		Height:   b.Height(),
		BarWidth: b.Width(),
		XAxis: chart.Style{
			Show: true,
		},
		YAxis: chart.YAxis{
			Style: chart.Style{
				Show: true,
			},
		},
		Bars: valuesToChart,
	}

	graph.Render(chart.PNG, w)
}
