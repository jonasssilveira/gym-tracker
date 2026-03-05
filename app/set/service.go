package set

import (
	"errors"
	"gym-tracker/infra/database/cache"
)

type Service struct {
	genericCache cache.GenericCache
	repo         SetsRepository
}

type SetCache interface {
	FinalizeSet(seriesID string) error
}

func NewService(genericCache cache.GenericCache, repo SetsRepository) Service {
	return Service{
		genericCache,
		repo,
	}
}

func (ss Service) FinalizeSet(seriesID uint64) error {
	sets, err := ss.genericCache.Get(seriesID)
	if err != nil {
		return err
	}
	if sets == nil {
		return errors.New("set service is empty")
	}

	for _, set := range sets {
		s := set.(SetDTO)
		_, err = ss.repo.CreateSet(s.ToEntity())
		if err != nil {
			return err
		}
	}
	return nil
}

func (ss Service) AddSet(set SetDTO) error {
	if invalid := set.Validate(); invalid != nil {
		return errors.New("set is invalid")
	}
	_, err := ss.genericCache.Set(set.SeriesID, set)
	return err
}

func (ss Service) GetALlSetsFromSerie(id uint64) ([]Set, error) {
	return ss.repo.GetALlSetsFromSerie(id)
}
