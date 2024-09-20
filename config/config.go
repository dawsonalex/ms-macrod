// Package config provides utilities and type
package config

import (
	"flag"
	"github.com/sirupsen/logrus"
	"github.com/wlevene/ini"
	"os"
)

var flagPath = flag.String("c", ".config.local.ini", "Path to config INI")

func init() {
	flag.Parse()
}

// C contains all available config types for the service.
type C struct {
	Server Server `ini:"http"`
	Log    Log    `ini:"log"`
}

var Default = &C{
	Server: Server{
		Host: "localhost",
		Port: "8432",
	},
	Log: Log{
		Level: logrus.InfoLevel,
	},
}

// ParseFile parses an INI file and returns a constructed config.C. path can be any INI file,
// but config.FlagPath is provided for convenience.
func ParseFile(path string) (*C, error) {
	confBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var conf C
	err = ini.Unmarshal(confBytes, &conf)
	if err != nil {
		return nil, err
	}

	return &conf, nil
}

// FlagPath returns the path to the config file provided to the `-c` flag.
func FlagPath() string {
	return *flagPath
}
