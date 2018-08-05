package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/D-Technologies/go-tokentracker/application"

	"github.com/D-Technologies/go-tokentracker/di"
	"github.com/D-Technologies/go-tokentracker/lib/config"
)

func main() {
	config.Setup()

	app := application.NewApp(
		"0x8f3f6c17ae5fd5B09fD070216995B6ebF3B224dD",
		di.InjectBlockNumRepository(),
		http.DefaultClient,
		di.InjectEthClient(),
		di.InjectSQL(),
	)

	if err := app.Do(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
	}
}
