package main

import (
	"flag"
	"fmt"
	"os"

	fzg "github.com/krakozaure/fzg/internal"
)

func main() {
	fzg.InitFlags()

	config, err := fzg.LoadConfig()
	if err != nil {
		handleError(err)
	}

	if fzg.CommandFlag != "" && fzg.OptionsFlag != "" {
		err = fzg.UseCommandAndOptions(config, fzg.CommandFlag, fzg.OptionsFlag)
	} else if fzg.CommandFlag != "" {
		err = fzg.UseCommand(config, fzg.CommandFlag)
	} else if fzg.OptionsFlag != "" {
		err = fzg.UseOptions(config, fzg.OptionsFlag)
	} else if fzg.ProfileFlag != "" {
		err = fzg.UseProfile(config, fzg.ProfileFlag)
	} else {
		flag.Usage()
		os.Exit(1)
	}
	if err != nil {
		handleError(err)
	}
}

func handleError(err error) {
	if !fzg.QuietFlag {
		fmt.Fprintln(os.Stderr, err)
	}
	os.Exit(1)
}
