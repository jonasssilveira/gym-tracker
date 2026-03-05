package set

import "gorm.io/gorm"

type SetsRepository struct {
	db gorm.DB
}

type Repository interface {
	GetALlSetsFromSerie(serieID uint64) ([]Set, error)
	GetSetByID(id int) (Set, error)
	CreateSet(sets Set) (Set, error)
}

func NewSetsRepository(db *gorm.DB) SetsRepository {
	return SetsRepository{db: *db}
}

func (r SetsRepository) GetALlSetsFromSerie(serieID uint64) ([]Set, error) {
	var sets []Set
	if err := r.db.Find(&sets).Where("series_id = ?", serieID).Error; err != nil {
		return nil, err
	}
	return sets, nil
}

func (r SetsRepository) GetSetByID(id int) (Set, error) {
	var sets Set
	if err := r.db.First(&sets, id).Error; err != nil {
		return Set{}, err
	}
	return sets, nil
}

func (r SetsRepository) CreateSet(sets Set) (Set, error) {
	if err := r.db.Create(&sets).Error; err != nil {
		return Set{}, err
	}
	return sets, nil
}
