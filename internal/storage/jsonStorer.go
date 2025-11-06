package storage

import (
	"errors"
	"fmt"

	"github.com/FerreiraLuciano/tp-go/internal/helper"
)

type JsonStore struct {
	nextID   int
	filePath string
}

func NewJsonStore(filePath string) *JsonStore {
	return &JsonStore{
		nextID:   1,
		filePath: filePath,
	}
}

func (j *JsonStore) Add(contact *Contact) error {

	targets, err := j.GetAll()

	id := findNextID(targets)

	newContact := helper.InputTarget{
		ID:    id,
		Name:  contact.Name,
		Email: contact.Email,
	}

	j.nextID = id

	var result []helper.InputTarget

	for _, target := range targets {
		result = append(result, convertToTarget(target))
	}

	result = append(result, newContact)

	err = helper.SaveTargetsToFile(j.filePath, result)
	if err != nil {
		return fmt.Errorf("Error adding new contact : %v\n", err)
	} else {
		return nil
	}
}

func (j *JsonStore) GetAll() ([]*Contact, error) {

	existingTargets, err := helper.LoadTargetsFromFile(j.filePath)
	if err != nil {
		return nil, fmt.Errorf("Error loading existing contacts : %v\n", err)
	}

	contacts := make([]*Contact, 0, len(existingTargets))
	for _, target := range existingTargets {
		contacts = append(contacts, convertToContact(target))
	}

	return contacts, nil
}

func (j *JsonStore) GetById(ID int) (*Contact, error) {

	existingTargets, err := helper.LoadTargetsFromFile(j.filePath)
	if err != nil {
		return nil, fmt.Errorf("Error loading existing contacts : %v\n", err)
	}

	contacts := make([]*Contact, 0, len(existingTargets))
	for _, target := range existingTargets {
		contacts = append(contacts, convertToContact(target))
	}

	for _, contact := range contacts {
		if contact.ID == ID {
			return contact, nil
		}
	}

	return nil, errors.New("Contact not found")
}

func (j *JsonStore) Update(ID int, name string, email string) error {

	contacts, err := j.GetAll()
	if err != nil {
		return err
	}

	target, err := j.GetById(ID)
	if err != nil {
		return err
	}

	target.Name = name
	target.Email = email

	for key, contact := range contacts {
		if contact.ID == ID {
			contacts[key] = target
		}
	}

	targets := make([]helper.InputTarget, 0, len(contacts))
	for _, contact := range contacts {
		targets = append(targets, convertToTarget(contact))
	}

	err = helper.SaveTargetsToFile(j.filePath, targets)
	if err != nil {
		return err
	}

	return nil
}

func (j *JsonStore) Delete(ID int) error {

	contacts, err := j.GetAll()
	if err != nil {
		return err
	}

	_, err = j.GetById(ID)
	if err != nil {
		return err
	}

	targets := make([]helper.InputTarget, 0)
	for _, contact := range contacts {
		if contact.ID != ID {
			targets = append(targets, convertToTarget(contact))
		}
	}

	err = helper.SaveTargetsToFile(j.filePath, targets)
	if err != nil {
		return err
	}

	return nil
}

func convertToContact(target helper.InputTarget) *Contact {
	return &Contact{
		ID:    target.ID,
		Name:  target.Name,
		Email: target.Email,
	}
}

func convertToTarget(contact *Contact) helper.InputTarget {
	return helper.InputTarget{
		ID:    contact.ID,
		Name:  contact.Name,
		Email: contact.Email,
	}
}

func findNextID(targets []*Contact) int {
	var id int

	for i, e := range targets {
		if i == 0 || e.ID > id {
			id = e.ID
		}
	}

	return id + 1
}
