package cmd

import (
	"fmt"
	"strconv"

	"github.com/FerreiraLuciano/tp-go/internal/helper"
	"github.com/FerreiraLuciano/tp-go/internal/storage"
	"github.com/spf13/cobra"
)

var (
	deleteFilePath string
	deleteId       string
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes an existing contact from the json file.",
	Long:  `The command 'delete' deletes an existing contact (ID, name, email) from the json file. If the contact is not in the file, it does nothing.`,
	Run: func(cmd *cobra.Command, args []string) {
		if deleteFilePath == "" || deleteId == "" {
			fmt.Println("Error: all the flags (--filePath, --id) are mandatory.")
			return
		}

		contactId, err := strconv.Atoi(deleteId)
		if err != nil {
			fmt.Printf("ID must be an int : %v\n", err)
			return
		}

		existingTargets, err := helper.LoadTargetsFromFile(deleteFilePath)

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

		targets := make([]helper.InputTarget, 0)
		for index, contact := range contacts {
			if existingTargets[index].ID != contactId {
				targets = append(targets, storage.ConvertToTarget(contact))
			}
		}

		err = helper.SaveTargetsToFile(deleteFilePath, targets)
		if err != nil {
			fmt.Printf("Error while updating contact : %v\n", err)
		} else {
			fmt.Printf("✅ Contact '%s' successfuly deleted to file '%s'.\n", contactFound.Name, deleteFilePath)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().StringVarP(&deleteFilePath, "file", "f", "", "Path to file to delete the contact to")
	deleteCmd.Flags().StringVarP(&deleteId, "id", "i", "", "Contact's id")

	deleteCmd.MarkFlagRequired("file")
	deleteCmd.MarkFlagRequired("id")

}
