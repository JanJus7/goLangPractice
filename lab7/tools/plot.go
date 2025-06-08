package tools

import (
	"fmt"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func GeneratePlot(dates []string, values []float64, title, yLabel, filename string) error {
	p := plot.New()
	p.Title.Text = title
	p.X.Label.Text = "Czas"
	p.Y.Label.Text = yLabel

	pts := make(plotter.XYs, len(dates))
	var isHourly bool

	for i := range dates {
		var t time.Time
		var err error

		if len(dates[i]) > 10 {
			t, err = time.Parse("2006-01-02T15:04", dates[i])
			isHourly = true
		} else {
			t, err = time.Parse("2006-01-02", dates[i])
		}

		if err != nil {
			return fmt.Errorf("date parse error: %q: %v", dates[i], err)
		}

		pts[i].X = float64(t.Unix())
		pts[i].Y = values[i]
	}

	if err := plotutil.AddLinePoints(p, title, pts); err != nil {
		return fmt.Errorf("plot points error: %v", err)
	}

	if isHourly {
		p.X.Tick.Marker = plot.TimeTicks{Format: "15:04"}
	} else {
		p.X.Tick.Marker = plot.TimeTicks{Format: "2006-01-02"}
	}

	if err := p.Save(8*vg.Inch, 4*vg.Inch, filename); err != nil {
		return fmt.Errorf("saving plot error: %v", err)
	}

	return nil
}
