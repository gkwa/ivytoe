package main

import (
	"github.com/taylormonacelli/ivytoe"
	"os"
)

func main() {
	status := ivytoe.Main()
	os.Exit(status)
}
