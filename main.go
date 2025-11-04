package main

import (
	"github.com/FerreiraLuciano/tp-go/internal/app"
	"github.com/FerreiraLuciano/tp-go/internal/storage"
)

func main() {

	store := storage.NewMemoryStore()

	app.Crm(store)
}
