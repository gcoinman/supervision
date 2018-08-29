package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/D-Technologies/supervision/application"

	"github.com/D-Technologies/supervision/di"
	"github.com/D-Technologies/supervision/lib/config"
)

func main() {
	config.Setup()

	app := application.NewApp(
		12985,
		"0x6176e9ec8ab713e3ab4ca415d25f57eea52e3cd6",
		"0xf007ebf754666e2216399fcbd59243845d8ac555",
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
