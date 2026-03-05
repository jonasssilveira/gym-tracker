package series

import (
	"errors"
	"fmt"
	"gym-tracker/app/set"
	"time"
)

type SeriesDTO struct {
	ID        uint64       `json:"id"`
	Name      string       `json:"name"`
	UserID    uint64       `json:"user_id"`
	TotalTime int          `json:"total_time"`
	Date      string       `json:"date"`
	Sets      []set.SetDTO `json:"sets"`
}

func NewSeriesDTO(name string, userID uint64) *SeriesDTO {
	now := time.Now()
	return &SeriesDTO{
		Name:   name,
		UserID: userID,
		Date:   fmt.Sprintf("%d-%d-%d", now.Year(), now.Month(), now.Day()),
		Sets:   nil,
	}
}

func (d SeriesDTO) Validate() error {
	if d.Name == "" {
		return errors.New("name is required")
	}
	return nil
}

func FromEntity(series Series) SeriesDTO {
	return SeriesDTO{
		ID:        series.ID,
		Name:      series.Name,
		TotalTime: series.TotalTime,
		Date:      fmt.Sprintf("%d-%d-%d", series.DateCreated.Year(), series.DateCreated.Month(), series.DateCreated.Day()),
		Sets:      set.FromSets(series.Sets),
	}
}

func (sDTO SeriesDTO) ToEntity() Series {
	return Series{
		ID:          sDTO.ID,
		Name:        sDTO.Name,
		TotalTime:   sDTO.TotalTime,
		DateCreated: time.Now(),
	}
}
