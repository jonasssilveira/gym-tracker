package set

import (
	"errors"
)

type SetDTO struct {
	ID       uint64 `json:"id"`
	SeriesID uint64 `json:"series_id"`
	Weight   int    `json:"weight"`
	Time     int    `json:"time"`
	Reps     int    `json:"reps"`
	RestTime int    `json:"real_time"`
}

func NewDTO(serieID uint64, weight, restTime, reps int) SetDTO {
	return SetDTO{
		SeriesID: serieID,
		Weight:   weight,
		RestTime: restTime,
		Reps:     reps,
	}
}

func FromSets(sets []Set) []SetDTO {
	var setDTOs []SetDTO
	for _, set := range sets {
		setDTOs = append(setDTOs, FromEntity(set))
	}
	return setDTOs
}

func FromEntity(set Set) SetDTO {
	return SetDTO{
		ID:       set.ID,
		Weight:   set.Weight,
		Time:     set.Time,
		RestTime: set.RestTime,
		Reps:     set.Reps,
	}
}

func (sDTO SetDTO) ToEntity() Set {
	return Set{
		SeriesID: sDTO.SeriesID,
		Weight:   sDTO.Weight,
		Time:     sDTO.Time,
		RestTime: sDTO.RestTime,
		Reps:     sDTO.Reps,
	}
}

func (sDTO SetDTO) Validate() error {
	if sDTO.Weight < 0 {
		return errors.New("weight cannot be 0 or negative")
	}
	if sDTO.RestTime < 0 {
		return errors.New("rest_time cannot be 0 or negative")
	}
	if sDTO.SeriesID < 0 {
		return errors.New("series_id cannot be 0 or negative")
	}
	return nil
}
