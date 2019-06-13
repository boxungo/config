package config

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ghodss/yaml"
)

type config struct {
	Name       string `json:"name"`
	configFile string
	flagSet    *flag.FlagSet
}

var usageline = `Usage:

	server [flags]
	  Start an server.

	server --config-file
	  Path to the server configuration file.
	`

func usage() {
	fmt.Fprintln(os.Stdout, usageline)
}

// NewConfig New Config
func NewConfig() *config {
	cfg := &config{}

	f := flag.NewFlagSet("name", flag.ContinueOnError)

	f.Usage = usage

	f.StringVar(&cfg.configFile, "c", "", "Path to the server configuration file")
	f.StringVar(&cfg.configFile, "config-file", "", "Path to the server configuration file")
	f.StringVar(&cfg.Name, "name", "", "Path to the server configuration file")

	if cfg.configFile != "" {
		err := cfg.configFromFile(cfg.configFile)
		if err != nil {
			fmt.Println("Load config file error, file: ", cfg.configFile, " error: ", err.Error())
		} else {
			fmt.Println("Load config file success, file: ", cfg.configFile)
		}
	}
	cfg.flagSet = f

	return cfg
}

func (cfg *config) configFromFile(configFile string) error {
	b, rerr := ioutil.ReadFile(configFile)
	if rerr != nil {
		return rerr
	}
	if yerr := yaml.Unmarshal(b, &cfg); yerr != nil {
		return yerr
	}
	return nil
}

func (cfg *config) parse(arguments []string) error {
	err := cfg.flagSet.Parse(arguments)

	switch err {
	case nil:
	case flag.ErrHelp:
		fmt.Println(usageline)
		os.Exit(0)
	default:
		os.Exit(2)
	}

	return err
}
