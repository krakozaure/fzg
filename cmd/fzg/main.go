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
		if !fzg.QuietFlag {
			fmt.Fprintln(os.Stderr, err)
		}
		os.Exit(1)
	}

	conf, text := "", ""
	if fzg.CommandFlag != "" && fzg.OptionsFlag != "" {
		conf = fmt.Sprintf("command '%s' or options '%s'", fzg.CommandFlag, fzg.OptionsFlag)
		text = fzg.ParseConfig(
			config.Commands[fzg.CommandFlag],
			config.Options[fzg.OptionsFlag],
		)
	} else if fzg.CommandFlag != "" {
		conf = fmt.Sprintf("command '%s'", fzg.CommandFlag)
		text = fzg.ParseConfig(
			config.Commands[fzg.CommandFlag],
			nil,
		)
	} else if fzg.OptionsFlag != "" {
		conf = fmt.Sprintf("options '%s'", fzg.OptionsFlag)
		text = fzg.ParseConfig(
			nil,
			config.Options[fzg.OptionsFlag],
		)
	} else if fzg.ProfileFlag != "" {
		conf = fmt.Sprintf("profile '%s'", fzg.ProfileFlag)
		text = fzg.ParseConfig(
			config.Profiles[fzg.ProfileFlag].Command,
			config.Profiles[fzg.ProfileFlag].Options,
		)
	} else {
		flag.Usage()
		os.Exit(1)
	}

	if text == "" {
		if !fzg.QuietFlag {
			fmt.Fprintf(os.Stderr, "Missing or invalid configuration %s\n", conf)
		}
		os.Exit(1)
	}
	fmt.Print(text)
}
