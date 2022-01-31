package fzg

import (
	"fmt"
	"strings"
)

func UseCommandAndOptions(config Config, commandKey string, optionsKey string) error {
	commandConf, ok := config.Commands[commandKey]
	if !ok {
		return fmt.Errorf("Unable to find a command configuration for '%s'", commandKey)
	}
	optionsConf, ok := config.Options[optionsKey]
	if !ok {
		return fmt.Errorf("Unable to find options configuration for '%s'", optionsKey)
	}
	return printCommandAndOptions(commandConf, optionsConf)
}

func UseCommand(config Config, key string) error {
	commandConf, ok := config.Commands[key]
	if !ok {
		return fmt.Errorf("Unable to find a command configuration for '%s'", key)
	}
	return printCommand(commandConf)
}

func UseOptions(config Config, key string) error {
	optionsConf, ok := config.Options[key]
	if !ok {
		return fmt.Errorf("Unable to find options configuration for '%s'", key)
	}
	return printOptions(optionsConf)
}

func UseProfile(config Config, key string) error {
	profileConf, ok := config.Profiles[key]
	if !ok {
		return fmt.Errorf("Unable to find a profile configuration for '%s'", key)
	}
	return printCommandAndOptions(profileConf.Command, profileConf.Options)
}

func printCommandAndOptions(commandConf Command, optionsConf Options) error {
	var err error
	err = printCommand(commandConf)
	if err != nil {
		return err
	}
	err = printOptions(optionsConf)
	if err != nil {
		return err
	}
	return err
}

func printCommand(commandConf Command) error {
	command := ""
	if commandConf != nil {
		command = parseCommand(commandConf)
	}
	if len(command) == 0 {
		return fmt.Errorf("Empty variable for the command")
	}

	if RawFlag {
		fmt.Printf("%s\n", command)
	} else {
		fmt.Printf("export %s=%q\n", "FZF_DEFAULT_COMMAND", command)
	}
	return nil
}

func printOptions(optionsConf Options) error {
	options := parseOptions(optionsConf)
	if len(options) == 0 {
		return fmt.Errorf("Empty variable for the options")
	}

	if RawFlag {
		fmt.Printf("%s\n", strings.Join(options, " "))
	} else {
		fmt.Printf("export %s=%q\n", "FZF_DEFAULT_OPTS", strings.Join(options, " "))
	}
	return nil
}
