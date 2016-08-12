package graph

import (
	"io"
	"time"

	"github.com/wcharczuk/go-chart"
)

const Time_YMDHIS = "2006-01-02T15:04:05"	

type TimeSeriesService struct {

}

type TimeSeries struct {
	height int 
	width int
	values []*Value
}

func (p *TimeSeries) Height() int {
	return p.height
}

func (p *TimeSeries) Width() int {
	return p.width
}

func (p *TimeSeries) Values() []*Value {
	return p.values
}

func NewTimeSeriesGraph(h int, w int, vs []float64, ls []string) *TimeSeries {
	numberOfValues := len(vs)
	values := make([]*Value, numberOfValues)
	for i := 0; i < numberOfValues; i++ {
		values[i] = &Value{
			Value: vs[i],
			Label: ls[i],
		}
	}	
	
	return &TimeSeries{
		height: h,
		width: w,
		values: values,
	}
}

func (s *TimeSeriesService) Build(p *TimeSeries, w io.Writer) {
	values := p.Values()
	valuesToChart := make([]float64, len(values))
	times := make([]time.Time, len(values))
	for i := 0; i < len(values); i++ {
		valuesToChart[i] = values[i].Value
		times[i], _ = time.Parse(Time_YMDHIS, values[i].Label)	
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
		},
		Series: []chart.Series{
			chart.TimeSeries{
				XValues: times,
				YValues: valuesToChart,
			},
		},
	}

	graph.Render(chart.PNG, w)
}