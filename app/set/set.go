package set

type Set struct {
	ID       uint64 `gorm:"primarykey"`
	SeriesID uint64 `gorm:"not null"`
	Weight   int    `gorm:"not null"`
	Time     int    `gorm:"not null"`
	RestTime int    `gorm:"not null"`
	Reps     int    `gorm:"not null"`
}
