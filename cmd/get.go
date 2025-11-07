package cmd

import (
	"fmt"
	"strconv"

	"github.com/FerreiraLuciano/tp-go/internal/helper"
	"github.com/FerreiraLuciano/tp-go/internal/storage"
	"github.com/spf13/cobra"
)

var (
	getFilePath  string
	getIdContact string
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Gets an existing contact from the json file.",
	Long:  `The command 'get' gets an existing contact (ID, name, email) from the json file. If the contact is not in the file, it does nothing.`,
	Run: func(cmd *cobra.Command, args []string) {
		if getFilePath == "" || getIdContact == "" {
			fmt.Println("Error: all the flags (--filePath, --id) are mandatory.")
			return
		}

		contactId, err := strconv.Atoi(getIdContact)
		if err != nil {
			fmt.Printf("ID must be an int : %v\n", err)
			return
		}

		existingTargets, err := helper.LoadTargetsFromFile(getFilePath)

		var contactFound *storage.Contact
		for _, target := range existingTargets {
			if target.ID == contactId {
				contactFound = storage.ConvertToContact(target)
				break
			}
		}

		if err != nil {
			fmt.Printf("Error while trying to load existing contacts : %v\n", err)
			return
		}

		if contactFound == nil {
			fmt.Printf("Desired contact does not exist.")
			return
		} else {
			fmt.Printf("ID: %d, Name: %s, Email: %s\n", contactFound.ID, contactFound.Name, contactFound.Email)
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	getCmd.Flags().StringVarP(&getFilePath, "file", "f", "", "Path to file to get the contact to")
	getCmd.Flags().StringVarP(&getIdContact, "id", "i", "", "Contact's id")

	getCmd.MarkFlagRequired("file")
	getCmd.MarkFlagRequired("id")

}
