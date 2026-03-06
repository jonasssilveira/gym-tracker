package series

import (
	"gorm.io/gorm"
)

type SeriesRepository struct {
	db gorm.DB
}

type Repository interface {
	GetSeriesByID(id uint64) ([]Series, error)
	GetSeriesByChatID(chatID int64) (Series, error)
	GetSeriesByName(name string) (Series, error)
	GetActualSeries() (Series, error)
	CreateSeries(series Series) (Series, error)
	FinalizeSerie(seriesID uint64, totalTime int) error
}

func NewSeriesRepository(db *gorm.DB) SeriesRepository {
	return SeriesRepository{db: *db}
}

func (r SeriesRepository) GetSeriesByChatID(chatID int64) ([]Series, error) {
	var series []Series
	if err := r.db.Preload("Sets").Find(&series, "chat_id = ?", chatID).Error; err != nil {
		return nil, err
	}
	return series, nil
}

func (r SeriesRepository) GetSeriesByID(id uint64) (Series, error) {
	var series Series
	if err := r.db.Preload("Set").First(&series, id).Error; err != nil {
		return Series{}, err
	}
	return series, nil
}

func (r SeriesRepository) GetSeriesByName(name string) (Series, error) {
	var series Series
	if err := r.db.Preload("Sets").First(&series, "name = ?", name).Error; err != nil {
		return Series{}, err
	}
	return series, nil
}

func (r SeriesRepository) GetActualSeries() (Series, error) {
	var series Series
	if err := r.db.Raw(
		"select * from series where finished = ? order by date_created limit 1", false,
	).Scan(&series).Error; err != nil {
		return Series{}, err
	}
	return series, nil
}

func (r SeriesRepository) CreateSeries(series Series) (Series, error) {
	if err := r.db.Create(&series).Error; err != nil {
		return Series{}, err
	}
	return series, nil
}

func (r SeriesRepository) FinalizeSerie(seriesID uint64, totalTime int) error {
	if err := r.db.Model(&Series{}).Where("id = ?", seriesID).Update("finished", true).Update("total_time", totalTime).Error; err != nil {
		return err
	}
	return nil
}
