package main

import (
	"errors"
	"flag"
	"fmt"
)

type Action string

const (
	Find   = "find"
	Delete = "delete"
	Create = "create"
)

func NewConfig() *Config {
	var cfg = Config{
		Folder:  "./tmp",
		File:    "",
		Action:  "find",
		Confirm: true,
	}

	flag.StringVar(&cfg.Folder, "folder", cfg.Folder, "The name of the `directory` to search")
	flag.StringVar(&cfg.File, "file", cfg.File, "Searched file `name`")
	flag.Func("action", "Performing action `name`: find/delete/create (default \"find\")", parseAction(&cfg))
	flag.BoolVar(&cfg.Confirm, "confirm", cfg.Confirm, "File deletion confirmation (for -action=delete)")
	flag.Parse()

	return &cfg
}

type Config struct {
	Folder  string
	File    string
	Action  Action
	Confirm bool
}

func (cfg *Config) String() string {
	return fmt.Sprint(
		"Action: ", cfg.Action, "\n",
		"Folder: ", cfg.Folder, "\n",
		"File:   ", cfg.File,
	)
}

func parseAction(cfg *Config) func(string) error {
	return func(act string) error {
		switch act {
		case "find":
			cfg.Action = Find
		case "delete":
			cfg.Action = Delete
		case "create":
			cfg.Action = Create
		default:
			return errors.New("unknown action")
		}

		return nil
	}
}
