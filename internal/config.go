package fzg

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

const (
	CONFIG_FILE_NAME = "fzg.yaml"
	CONFIG_FILE_PATH = "$HOME/.config/fzg/fzg.yaml"
)

type Config struct {
	Commands map[string]Command `yaml:"commands"`
	Options  map[string]Options `yaml:"options"`
	Profiles map[string]Profile `yaml:"profiles"`
}

type Profile struct {
	Command Command
	Options Options
}

func LoadConfig() (Config, error) {
	var conf Config

	confFile, err := configFile()
	if err != nil {
		return conf, err
	}

	confRaw, err := ioutil.ReadFile(confFile)
	if err != nil {
		return conf, nil
	}

	err = yaml.Unmarshal(confRaw, &conf)
	return conf, err
}

func ParseConfig(commandConf Command, optionsConf Options) string {
	command := parseCommand(commandConf)
	options := parseOptions(optionsConf)
	output := ""
	if !RawFlag {
		if command != "" {
			output += fmt.Sprintf("export %s=%q\n", "FZF_DEFAULT_COMMAND", command)
		}
		if options != "" {
			output += fmt.Sprintf("export %s=%q\n", "FZF_DEFAULT_OPTS", options)
		}
	} else {
		if command != "" {
			output += fmt.Sprintf("%s\n", command)
		}
		if options != "" {
			output += fmt.Sprintf("%s\n", options)
		}
	}
	return output
}

func configFile() (string, error) {
	var confFile string

	envConfigFile := os.Getenv("FZG_CONFIG_FILE")
	xdgConfigFile := os.ExpandEnv(CONFIG_FILE_PATH)

	if envConfigFile != "" && isFile(envConfigFile) {
		return envConfigFile, nil
	} else if isFile(xdgConfigFile) {
		return xdgConfigFile, nil
	} else if isFile(CONFIG_FILE_NAME) {
		return CONFIG_FILE_NAME, nil
	} else {
		return confFile, fmt.Errorf("Unable to find the configuration file")
	}
}

func isFile(path string) bool {
	info, err := os.Stat(path)
	if err != nil || os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
