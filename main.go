package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/D-Technologies/go-tokentracker/application"

	"github.com/D-Technologies/go-tokentracker/di"
	"github.com/D-Technologies/go-tokentracker/lib/config"
)

func main() {
	config.Setup()

	app := application.NewApp(
		"0x232d9d6451bf3a2600cee974262e46aab02d4e0f",
		"0x8f3f6c17ae5fd5B09fD070216995B6ebF3B224dD",
		di.InjectBlockNumRepository(),
		di.InjectReceivedTransactionRepository(),
		di.InjectConfirmedTransactionRepository(),
		http.DefaultClient,
		di.InjectEthClient(),
		di.InjectSQL(),
	)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		for {
			if err := app.Do(); err != nil {
				fmt.Fprintln(os.Stderr, "Error:", err)
			}
			time.Sleep(15 * time.Second)
		}
	}()

	wg.Wait()
}
