package main

import "github.com/guptarohit/asciigraph"

const historySize = 200

func appendHistory(history []float64, value float64) []float64 {
	history = append(history, value)
	if len(history) > historySize {
		history = history[1:]
	}
	return history
}

func smooth(history []float64, window int) []float64 {
	if len(history) < window {
		return history
	}
	smoothed := make([]float64, len(history))
	for i := range history {
		start := i - window + 1
		if start < 0 {
			start = 0
		}
		sum := 0.0
		for _, v := range history[start : i+1] {
			sum += v
		}
		smoothed[i] = sum / float64(i-start+1)
	}
	return smoothed
}

func renderCPUGraph(history []float64, width, height int) string {
	if len(history) == 0 {
		return ""
	}
	return asciigraph.Plot(smooth(history, 5),
		asciigraph.Height(height),
		asciigraph.Width(width),
		asciigraph.LowerBound(0),
		asciigraph.UpperBound(100),
		asciigraph.SeriesColors(asciigraph.Green),
	)
}
func renderDownloadGraph(history []float64, width, height int) string {
	if len(history) == 0 {
		return ""
	}
	return asciigraph.Plot(smooth(history, 5),
		asciigraph.Height(height),
		asciigraph.Width(width),
		asciigraph.LowerBound(0),
		asciigraph.SeriesColors(asciigraph.Cyan),
	)
}

func renderUploadGraph(history []float64, width, height int) string {
	if len(history) == 0 {
		return ""
	}
	return asciigraph.Plot(smooth(history, 5),
		asciigraph.Height(height),
		asciigraph.Width(width),
		asciigraph.LowerBound(0),
		asciigraph.SeriesColors(asciigraph.Magenta),
	)
}
