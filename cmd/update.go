package cmd

import (
	"fmt"
	"strconv"

	"github.com/FerreiraLuciano/tp-go/internal/helper"
	"github.com/FerreiraLuciano/tp-go/internal/storage"
	"github.com/spf13/cobra"
)

var (
	updateFilePath string
	updateId       string
	updateName     string
	updateEmail    string
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an existing contact to the desired json file.",
	Long:  `The command 'update' updates an existing contact (name, email) in the desired json file. If the contact does not exist, it returns an error.`,
	Run: func(cmd *cobra.Command, args []string) {
		if updateFilePath == "" || updateId == "" || updateName == "" || updateEmail == "" {
			fmt.Println("Error: all the flags (--file, --id, --name, --email) are mandatory.")
			return
		}

		contactId, err := strconv.Atoi(updateId)
		if err != nil {
			fmt.Printf("ID must be an int : %v\n", err)
			return
		}

		existingTargets, err := helper.LoadTargetsFromFile(updateFilePath)

		var contacts []*storage.Contact
		var contactFound *storage.Contact
		for _, target := range existingTargets {
			contacts = append(contacts, storage.ConvertToContact(target))

			if target.ID == contactId {
				contactFound = storage.ConvertToContact(target)
			}
		}

		if contactFound == nil {
			fmt.Printf("Desired contact does not exist.")
			return
		}

		if err != nil {
			fmt.Printf("Error while trying to load existing contacts : %v\n", err)
			return

		}

		for index := range contacts {
			if existingTargets[index].ID == contactId {
				existingTargets[index].Name = updateName
				existingTargets[index].Email = updateEmail
				break
			}
		}

		err = helper.SaveTargetsToFile(updateFilePath, existingTargets)
		if err != nil {
			fmt.Printf("Error while updating contact : %v\n", err)
		} else {
			fmt.Printf("✅ Contact '%s' successfuly updated to file '%s'.\n", contactFound.Name, updateFilePath)
		}

	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	updateCmd.Flags().StringVarP(&updateFilePath, "file", "f", "", "Path of the file to update the contact to")
	updateCmd.Flags().StringVarP(&updateId, "id", "i", "", "Contact's id")
	updateCmd.Flags().StringVarP(&updateName, "name", "n", "", "Contact's name")
	updateCmd.Flags().StringVarP(&updateEmail, "email", "u", "", "Contact's email")

	updateCmd.MarkFlagRequired("file")
	updateCmd.MarkFlagRequired("id")
	updateCmd.MarkFlagRequired("name")
	updateCmd.MarkFlagRequired("email")
}
