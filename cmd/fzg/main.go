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

	err_text, out_text := "", ""
	if fzg.CommandFlag != "" && fzg.OptionsFlag != "" {
		err_text = fmt.Sprintf(
			"Invalid or missing configuration for '%s' or '%s'",
			fzg.CommandFlag, fzg.OptionsFlag,
		)
		out_text = fzg.ParseConfig(
			config.Commands[fzg.CommandFlag],
			config.Options[fzg.OptionsFlag],
		)
	} else if fzg.CommandFlag != "" {
		err_text = fmt.Sprintf(
			"Invalid or missing command configuration for '%s'",
			fzg.CommandFlag,
		)
		out_text = fzg.ParseConfig(
			config.Commands[fzg.CommandFlag],
			nil,
		)
	} else if fzg.OptionsFlag != "" {
		err_text = fmt.Sprintf(
			"Invalid or missing options configuration for '%s'",
			fzg.OptionsFlag,
		)
		out_text = fzg.ParseConfig(
			nil,
			config.Options[fzg.OptionsFlag],
		)
	} else if fzg.ProfileFlag != "" {
		err_text = fmt.Sprintf(
			"Invalid or missing profile configuration for '%s'",
			fzg.ProfileFlag,
		)
		out_text = fzg.ParseConfig(
			config.Profiles[fzg.ProfileFlag].Command,
			config.Profiles[fzg.ProfileFlag].Options,
		)
	} else {
		flag.Usage()
		os.Exit(1)
	}

	if out_text == "" {
		if !fzg.QuietFlag {
			fmt.Fprintf(os.Stderr, "%s\n", err_text)
		}
		os.Exit(1)
	}
	fmt.Print(out_text)
}
