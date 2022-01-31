package fzg

import (
	"flag"
	"fmt"
	"os"
)

var (
	QuietFlag bool
	RawFlag   bool

	CommandFlag string
	OptionsFlag string
	ProfileFlag string
)

func InitFlags() {
	flag.BoolVar(
		&QuietFlag,
		"q",
		false,
		"Fail without printing errors but with exit code > 0 (default: false)",
	)
	flag.BoolVar(
		&RawFlag,
		"r",
		false,
		"Print raw value without variable name or quoting (default: false)",
	)

	flag.StringVar(
		&CommandFlag,
		"c",
		"",
		"Configuration key to use for the command",
	)
	flag.StringVar(
		&OptionsFlag,
		"o",
		"",
		"Configuration key to use for the options",
	)
	flag.StringVar(
		&ProfileFlag,
		"p",
		"",
		"Configuration key to use for the profile (command+options)",
	)

	flag.Usage = func() {
		fmt.Fprintf(
			os.Stderr,
			`USAGE: %s [-q] [-r] [-c CMD -o OPTS | -c CMD | -o OPTS | -p PROFILE]

OPTIONS:
`,
			os.Args[0],
		)
		flag.PrintDefaults()
	}

	flag.Parse()
}
