package calculos

import "time"

type Calculos interface {
	calculate() Results
}

type Results struct {
	results map[time.Time]SerieVolume
}

type SerieVolume struct {
	serieName string
	volume    int
}
