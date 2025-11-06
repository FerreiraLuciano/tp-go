package cmd

import (
	"github.com/FerreiraLuciano/tp-go/internal/app"
	"github.com/FerreiraLuciano/tp-go/internal/storage"
	"github.com/spf13/cobra"
)

var cliCmd = &cobra.Command{
	Use:   "cli",
	Short: "This is a cli for a crm about Denver",
	Run: func(cmd *cobra.Command, args []string) {
		store := storage.NewMemoryStore()
		app.Crm(store)
	},
}

func init() {
	rootCmd.AddCommand(cliCmd)
}
