package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/FerreiraLuciano/tp-go/internal/config"
	"github.com/FerreiraLuciano/tp-go/internal/storage"
	"github.com/spf13/cobra"
)

var Cfg *config.Config

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
	Cfg, err = config.LoadConfig()

	if err != nil {
		log.Fatal("Error loading configuration:", err)
	}

	switch Cfg.Storage.Type {
	case "memory":
		store = storage.NewMemoryStore()
	case "json":
		store = storage.NewJsonStore(Cfg.Storage.Path)
	case "database":
		store, err = storage.NewGORMStore(Cfg.Storage.Path)
	default:
		log.Fatal("Unknown storage type in configuration")
	}

	if err != nil {
		log.Fatal("Error while establishing connection to the storage:", err)
	}
}
