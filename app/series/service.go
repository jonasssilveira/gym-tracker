package series

import (
	"log/slog"
	"time"
)

var logger = slog.Default()

type Service struct {
	repo SeriesRepository
}

type SerieService interface {
	FinalizeSerie(seriesID string) error
}

func NewService(repo SeriesRepository) Service {
	return Service{repo: repo}
}

func (ss Service) GetALlSeriesByChatID(userID int64) []Series {
	series, err := ss.repo.GetSeriesByChatID(userID)
	if err != nil {
		logger.Warn("Series Not Found")
		return []Series{}
	}
	return series
}

func (ss Service) FinalizeSerie() (uint64, error) {
	serie, err := ss.ActualSerie()
	now := time.Now()
	if err != nil {
		return 0, err
	}

	diffMin := min(serie.DateCreated.Minute(), now.Minute())
	diffMax := max(serie.DateCreated.Minute(), now.Minute())
	exerciceTemp := diffMax - diffMin
	if err = ss.repo.FinalizeSerie(serie.ID, exerciceTemp); err != nil {
		return 0, err
	}
	return serie.ID, err
}

func (ss Service) CreateSeries(entity Series) (Series, error) {
	return ss.repo.CreateSeries(entity)
}

func (ss Service) SerieByName(name string) (Series, error) {
	return ss.repo.GetSeriesByName(name)
}

func (ss Service) ActualSerie() (Series, error) {
	return ss.repo.GetActualSeries()
}
