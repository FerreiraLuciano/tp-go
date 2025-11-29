package storage

import (
	"errors"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type GORMStore struct {
	db *gorm.DB
}

func NewGORMStore(dbPath string) (*GORMStore, error) {
	db, err := gorm.Open(sqlite.Open("sqlite.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&Contact{})
	if err != nil {
		return nil, err
	}

	return &GORMStore{db: db}, nil
}

func (g *GORMStore) GetAll() ([]*Contact, error) {
	var contacts []*Contact
	err := g.db.Find(&contacts).Error
	return contacts, err
}

func (g *GORMStore) GetById(ID int) (*Contact, error) {
	var contact Contact
	result := g.db.First(&contact, ID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("Contact not found")
		}
		return nil, result.Error
	}
	return &contact, nil
}

func (g *GORMStore) Add(contact *Contact) error {
	return g.db.Create(contact).Error
}

func (g *GORMStore) Delete(ID int) error {
	_, err := g.GetById(ID)
	if err != nil {
		return err
	}

	result := g.db.Unscoped().Delete(&Contact{}, ID)
	return result.Error
}

func (g *GORMStore) Update(ID int, name string, email string) error {
	contact, err := g.GetById(ID)
	if err != nil {
		return err
	}

	if name != "" {
		contact.Name = name
	}
	if email != "" {
		contact.Email = email
	}

	result := g.db.Save(contact)
	return result.Error
}
