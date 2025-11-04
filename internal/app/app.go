package app

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

import . "github.com/FerreiraLuciano/tp-go/internal/storage"

type UserList map[int]*Contact

func (u *Contact) newContact(id int, name string, email string) *Contact {

	if "" == name || "" == email {
		return nil
	}

	return &Contact{id, name, email}
}

func (u *Contact) createContact(contacts map[int]*Contact) {
	contacts[len(contacts)+1] = u.newContact(u.ID, u.Name, u.Email)
}

func (contacts *UserList) removeContact(id int) (string, error) {

	_, ok := (*contacts)[id]
	if ok {
		delete(*contacts, id)
		return fmt.Sprintf("Le contact d'id %d a bien été supprimé", id), nil
	} else {
		return "", errors.New("le contact n'existe pas")
	}

}

func Crm() {

	contacts := make(map[int]*Contact)
	var user1 = new(Contact)
	user1 = user1.newContact(len(contacts)+1, "Le premier", "lepremier@hihi.cl")
	user1.createContact(contacts)

	var user2 = new(Contact)
	user2 = user2.newContact(len(contacts)+1, "Le deuxième", "ledeuxieme@hihi.cl")
	user2.createContact(contacts)

	flagName := flag.String("name", "", "help message for flag name")
	flagMail := flag.String("email", "", "help message for flag email")
	flag.Parse()

	fmt.Println(*flagName, *flagMail)

	if "" != *flagName && "" != *flagMail {
		flagUser := new(Contact)
		flagUser = flagUser.newContact(len(contacts)+1, *flagName, *flagMail)
		flagUser.createContact(contacts)
	}

	for {
		printChoices()

		choice := getUserInput()

		switch choice {

		case "1", "a", "A":
			addContact(contacts)

		case "2", "l", "L":
			listContacts(&contacts)

		case "3", "u", "U":
			updateContact(contacts)

		case "4", "d", "D":
			deleteContact(contacts)

		case "5", "q", "Q":
			fmt.Println("Bye !")
			return

		}
	}
}

func printChoices() {
	fmt.Println("")
	fmt.Println("Choose an option : ")
	fmt.Println("	1 - Add contact")
	fmt.Println("	2 - List contacts")
	fmt.Println("	3 - Update contact")
	fmt.Println("	4 - Delete contact")
	fmt.Println("	5 - Quit")
}

func addContact(contacts map[int]*Contact) {

	fmt.Println("")
	fmt.Print("Type the new user's email : ")

	email := getUserInput()

	fmt.Print("Type the new user's name : ")

	name := getUserInput()

	id := len(contacts) + 1

	var user = new(Contact)
	user = user.newContact(id, name, email)

	user.createContact(contacts)

	fmt.Println("\nContact added successfully with ID", id, "!")
}

func listContacts(contacts *UserList) {

	fmt.Println("")
	fmt.Println("Contacts : ")
	for _, value := range *contacts {
		fmt.Println("	- ID -", value.ID, ": {", "name :", value.Name, ", email :", value.Email, "}")
	}
}

func updateContact(contacts UserList) {

	fmt.Println("")
	fmt.Println("Choose a contact to edit by its' id : ")

	id, _ := strconv.Atoi(getUserInput())

	value, ok := contacts[id]
	if !ok {
		log.Fatalf("Contact not found : %s", id)
	}

	fmt.Print("Update this contact's email (leave blank to keep the current one) : ")

	email := getUserInput()

	fmt.Print("Update this contact's name (leave blank to keep the current one) : ")

	name, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		log.Fatalf("Error reading input : %s", err)
	}
	name = strings.TrimSpace(name)

	if name == "" {
		name = value.Name
	}
	if email == "" {
		email = value.Email
	}

	contacts[id] = &Contact{value.ID, name, email}

	fmt.Println("\nContact updated successfully !")
}

func deleteContact(contacts UserList) {

	fmt.Println("")
	fmt.Println("Choose a contact by its' id : ", contacts)

	id, _ := strconv.Atoi(getUserInput())

	_, err := contacts.removeContact(id)
	if err != nil {
		return
	}

	fmt.Println("\nContact deleted successfully !")
}

func getUserInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Error reading input : %s", err)
	}
	return strings.TrimSpace(input)
}
