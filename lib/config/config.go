package config

import (
	"github.com/pkg/errors"

	"github.com/kelseyhightower/envconfig"
)

// TokenTracker represents a struct for mysql config
type TokenTracker struct {
	DbUser     string `default:"root" equired:"true" split_words:"true"`
	DbPassword string `default:"twice" equired:"true" split_words:"true"`
	DbHost     string `default:"127.0.0.1" equired:"true" split_words:"true"`
}

// DefaultConfig is a default config
var DefaultConfig TokenTracker

// Setup sets up a config
func Setup() {
	if err := setConf("go-tokentracker"); err != nil {
		panic(err)
	}
}

func setConf(name string) error {
	var a interface{}
	switch name {
	case "go-tokentracker":
		a = &DefaultConfig
	default:
		return errors.New("type not found")
	}

	if err := envconfig.Process(name, a); err != nil {
		return err
	}

	return nil
}
