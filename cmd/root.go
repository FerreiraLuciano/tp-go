package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/FerreiraLuciano/tp-go/internal/storage"
	"github.com/spf13/cobra"
)

var store storage.Storer

var rootCmd = &cobra.Command{
	Use:   "crm",
	Short: "This crm is a simple tool to manage contacts",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func init() {
	var err error
	store, err = storage.NewGORMStore("contacts.db")
	if err != nil {
		log.Fatal("Error while establishing connection to the database:", err)
	}
}
