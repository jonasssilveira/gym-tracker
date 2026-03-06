package calculos

import (
	"gym-tracker/app/series"
	"time"
)

type TrainingVolumeLoad struct {
}

func (tvl TrainingVolumeLoad) Calculate(series []series.Series) Results {
	var results Results
	serieVolume := make(map[time.Time]SerieVolume)
	for _, serie := range series {
		var volume int
		for _, set := range serie.Sets {
			volume += set.Weight * set.Reps
		}
		serieVolume[serie.DateCreated] = SerieVolume{
			serie.Name,
			volume,
		}
	}
	results.results = serieVolume
	return results
}
