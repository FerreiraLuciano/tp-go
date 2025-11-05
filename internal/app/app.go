package app

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/FerreiraLuciano/tp-go/internal/storage"
)

func NewContact(name string, email string) *storage.Contact {
	if "" == name || "" == email {
		return nil
	}

	return &storage.Contact{Name: name, Email: email}
}

func Crm(store storage.Storer) {

	reader := bufio.NewReader(os.Stdin)

	for {
		printChoices()

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1", "a", "A":
			handleAddContact(reader, store)
		case "2", "l", "L":
			handleListContacts(store)
		case "3", "u", "U":
			handleUpdateContact(reader, store)
		case "4", "d", "D":
			handleDeleteContact(reader, store)
		case "5", "q", "Q":
			fmt.Println("Ciao !")
			return
		default:
			fmt.Println("Unavailable option.")
		}
	}

}

func handleAddContact(reader *bufio.Reader, store storage.Storer) {
	fmt.Print("Name: ")
	name, _ := reader.ReadString('\n')

	fmt.Print("Email: ")
	email, _ := reader.ReadString('\n')

	contact := NewContact(strings.TrimSpace(name), strings.TrimSpace(email))

	if err := store.Add(contact); err != nil {
		fmt.Printf("Erreur lors de l'ajout: %v\n", err)
		return
	}

	fmt.Println("Contact successfully added !")
}

func handleUpdateContact(reader *bufio.Reader, store storage.Storer) {
	fmt.Print("ID du contact à modifier: ")
	id, _ := reader.ReadString('\n')
	id = strings.TrimSpace(id)
	target, err := strconv.Atoi(id)
	if err != nil {
		fmt.Printf("Invalid ID : %v\n", err)
		return
	}

	fmt.Print("New name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("New email: ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	if err := store.Update(target, name, email); err != nil {
		fmt.Printf("Erreur: %v\n", err)
		return
	}

	fmt.Println("Contact updated successfully !")
}

func handleListContacts(store storage.Storer) {

	contacts, err := store.GetAll()

	if err != nil {
		fmt.Printf("Erreur: %v\n", err)
		return
	}

	for _, contact := range contacts {
		fmt.Printf("ID: %d, Name: %s, Email: %s\n", contact.ID, contact.Name, contact.Email)
	}
}

func handleDeleteContact(reader *bufio.Reader, store storage.Storer) {

	fmt.Print("ID du contact à supprimer: ")
	id, _ := reader.ReadString('\n')
	id = strings.TrimSpace(id)
	target, err := strconv.Atoi(id)
	if err != nil {
		fmt.Printf("Erreur: %v\n", err)
		return
	}

	if err := store.Delete(target); err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("Contact deleted successfully !")
}

func printChoices() {
	fmt.Println("\nChoose an option : ")
	fmt.Println("	1 - Add contact")
	fmt.Println("	2 - List contacts")
	fmt.Println("	3 - Update contact")
	fmt.Println("	4 - Delete contact")
	fmt.Println("	5 - Quit")
	fmt.Print("Please enter your choice : ")
}
