package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/martinlindhe/notify"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	// show a notification
	// notify.Notify("Aeacus SE", "Phocus Client", "You have gained points!", filepath.Join(dir, "src/icons/aeacus.png"))
	notify.Alert("app name", "alert", "some text", filepath.Join(dir, "src/icons/aeacus.png"))
}
