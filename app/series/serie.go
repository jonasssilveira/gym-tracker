package series

import (
	"gym-tracker/app/set"
	"time"
)

type Series struct {
	ID          uint64    `gorm:"primarykey"`
	UserID      uint64    `gorm:"not null"`
	Name        string    `gorm:"not null"`
	TotalTime   int       `gorm:"not null"`
	DateCreated time.Time `gorm:"not null"`
	Sets        []set.Set `gorm:"foreignKey:SeriesID;references:ID"`
}
