package duplicate

import (
	"errors"
	"flag"
	"fmt"
)

// The type for actions (find/delete/create).
type Action string

const (
	Find   = "find"   // Action for find duplicates.
	Delete = "delete" // Action for deleting duplicates.
	Create = "create" // Action for creating duplicates.
)

// NewConfig creates new config and returns pointer to it.
func NewConfig() *Config {
	var cfg = Config{
		Folder: "./tmp",
		File:   "",
		Action: "find",
	}

	flag.StringVar(&cfg.Folder, "folder", cfg.Folder, "The name of the `directory` to search")
	flag.StringVar(&cfg.File, "file", cfg.File, "Searched file `name`")
	flag.Func("action", "Performing action `name`: find/delete/create (default \"find\")", parseAction(&cfg))
	flag.Parse()

	return &cfg
}

// The struct for configurations.
type Config struct {
	Folder string // Path to derictory when will be find.
	File   string // Path to file which will be used like original.
	Action Action // Action for execution.
}

func (cfg *Config) String() string {
	return fmt.Sprint(
		"Action: ", cfg.Action, "\n",
		"Folder: ", cfg.Folder, "\n",
		"File:   ", cfg.File,
	)
}

// parseAction returns parser for 'action' flag.
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
