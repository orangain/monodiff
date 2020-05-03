package main

import (
	"fmt"
	"log"
	"os"

	"github.com/orangain/monodiff"
)

// These variables are set in build step
var (
	Version  = "unset"
	Revision = "unset"
)

func main() {
	app := monodiff.App
	app.Version = fmt.Sprintf("%s (%s)", Version, Revision)
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
