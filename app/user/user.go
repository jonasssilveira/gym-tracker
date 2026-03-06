package user

import "time"

type User struct {
	ID          uint64    `gorm:"primarykey"`
	FullName    string    `gorm:"not null"`
	ChatID      int64     `gorm:"not null"`
	DateCreated time.Time `gorm:"not null"`
	DateUpdated time.Time `gorm:"not null"`
}
