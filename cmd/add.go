package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/FerreiraLuciano/tp-go/internal/helper"
	"github.com/FerreiraLuciano/tp-go/internal/storage"
	"github.com/spf13/cobra"
)

var (
	addFilePath string
	addName     string
	addEmail    string
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a new contact to the desired json file.",
	Long:  `The command 'add' creates a new contact (name, email) in the desired json file. If the file does not exist, it will be created..`,
	Run: func(cmd *cobra.Command, args []string) {
		if addFilePath == "" || addName == "" || addEmail == "" {
			fmt.Println("Error: all the flags (--file, --name, --email) are mandatory.")
			return
		}

		existingTargets, err := helper.LoadTargetsFromFile(addFilePath)

		var contacts []*storage.Contact
		for _, target := range existingTargets {
			contacts = append(contacts, storage.ConvertToContact(target))
		}

		newTarget := helper.InputTarget{
			ID:    storage.FindNextID(contacts),
			Name:  addName,
			Email: addEmail,
		}

		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				existingTargets = []helper.InputTarget{}
				fmt.Printf("The file '%s' does not exist. It will be created.\n", addFilePath)
			} else {
				fmt.Printf("Error while trying to load existing contacts : %v\n", err)
				return
			}
		}

		existingTargets = append(existingTargets, newTarget)

		err = helper.SaveTargetsToFile(addFilePath, existingTargets)
		if err != nil {
			fmt.Printf("Error while creating new contatc : %v\n", err)
		} else {
			fmt.Printf("✅ Contact '%s' successfuly added to file '%s'.\n", newTarget.Name, addFilePath)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringVarP(&addFilePath, "file", "f", "", "Path of the file to add contact to")
	addCmd.Flags().StringVarP(&addName, "name", "n", "", "Contact's name")
	addCmd.Flags().StringVarP(&addEmail, "email", "u", "", "Contact's email")

	addCmd.MarkFlagRequired("file")
	addCmd.MarkFlagRequired("name")
	addCmd.MarkFlagRequired("email")
}
