package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type User struct {
	Name  string
	Email string
}

type UserList map[string]User

func main() {

	crm()
}

func crm() {

	contacts := make(UserList)
	contacts["1"] = User{"Le premier", "lepremier@hihi.cl"}
	contacts["2"] = User{"Le deuxième", "ledeuxieme@haha.cl"}

	for {
		printChoices()

		choice := getUserInput()

		switch choice {

		case "1", "a", "A":
			addContact(contacts)

		case "2", "l", "L":
			listContacts(contacts)

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

func addContact(contacts UserList) {

	fmt.Println("")
	fmt.Print("Type the new user's email : ")

	email := getUserInput()

	fmt.Print("Type the new user's name : ")

	name := getUserInput()

	id := strconv.Itoa(len(contacts) + 1)

	contacts[id] = User{name, email}

	fmt.Println("\nContact added successfully with ID", id, "!")
}

func listContacts(contacts UserList) {

	fmt.Println("")
	fmt.Println("Contacts : ")
	for key, value := range contacts {
		fmt.Println("	- ID -", key, ": {", "name :", value.Name, ", email :", value.Email, "}")
	}
}

func updateContact(contacts UserList) {

	fmt.Println("")
	fmt.Println("Choose a contact to edit by its' id : ")

	id := getUserInput()

	value, ok := contacts[id]
	if !ok {
		log.Fatalf("User not found : %s", id)
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

	contacts[id] = User{name, email}

	fmt.Println("\nContact updated successfully !")
}

func deleteContact(contacts UserList) {

	fmt.Println("")
	fmt.Println("Choose a contact by its' id : ", contacts)

	id := getUserInput()

	_, ok := contacts[id]
	if !ok {
		fmt.Println("Contact not found")
	}

	delete(contacts, id)

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
