package storage

type Contact struct {
	ID    int
	Name  string
	Email string
}

type Storer interface {
	Add(contact *Contact) error
	GetAll() ([]*Contact, error)
	GetById(ID int) (*Contact, error)
	Update(ID int, name string, email string) error
	Delete(ID int) error
}
