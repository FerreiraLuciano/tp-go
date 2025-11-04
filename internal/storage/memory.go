package storage

import "errors"

type MemoryStore struct {
	contacts map[int]*Contact
	nextID   int
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		contacts: make(map[int]*Contact),
		nextID:   1,
	}
}

func (ms *MemoryStore) Add(contact *Contact) error {

	contact.ID = ms.nextID
	ms.contacts[contact.ID] = contact
	ms.nextID++
	return nil
}

func (ms *MemoryStore) GetAll() ([]*Contact, error) {
	contacts := make([]*Contact, 0, len(ms.contacts))
	for _, contact := range ms.contacts {
		contacts = append(contacts, contact)
	}
	return contacts, nil
}

func (ms *MemoryStore) GetById(ID int) (*Contact, error) {
	contact, ok := ms.contacts[ID]
	if !ok {
		return nil, errors.New("Contact not found")
	}
	return contact, nil
}

func (ms *MemoryStore) Update(ID int, name string, email string) error {
	contact, ok := ms.contacts[ID]
	if ok {
		contact.Name = name
		contact.Email = email
	}
	return nil
}

func (ms *MemoryStore) Delete(ID int) error {
	_, ok := ms.contacts[ID]
	if !ok {
		return errors.New("Contact not found")
	}
	delete(ms.contacts, ID)
	return nil
}
