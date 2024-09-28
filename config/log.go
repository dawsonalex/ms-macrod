package config

import "github.com/sirupsen/logrus"

type Log struct {
	Level logrus.Level `ini:"level"`

	IncludeFields bool `ini:"include_fields"`
}
