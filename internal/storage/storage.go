package storage

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	ID    int    `gorm:"primaryKey"`
	Name  string `gorm:"type:varchar(20);not null"`
	Email string `gorm:"type:varchar(50);uniqueIndex;not null"`
}

type Storer interface {
	Add(contact *Contact) error
	GetAll() ([]*Contact, error)
	GetById(ID int) (*Contact, error)
	Update(ID int, name string, email string) error
	Delete(ID int) error
}
