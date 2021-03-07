// Package config implements function to connect to external server
//
// The NewConnection opens connection and return error if failed
//
// NewConnection() error
//
// The CloseConnection closes connection
//
// CloseConnection() error
//
// The ValidateConfiguration validate data that was set in configuration
//
// ValidateConfiguration() error
package config

import (
	"errors"
	"flag"
	"github.com/kelseyhightower/envconfig"
	"log"
	"regexp"
)

type Configuration struct {
	Debug     bool   `envconfig:"DEBUG" default:"false" required:"true"`
	Port      int    `envconfig:"PORT" default:"8080" required:"true"`
	DBURL     string `envconfig:"DB_URL" default:"postgres://user:password@localhost:5432/petstore?sslmode=disable" required:"true"`
	connected bool
}

//NewConnection set new connection
func (conf *Configuration) NewConnection() error {
	timeout := flag.Int("timeout", 30, "Timeout of connection in seconds")

	defer func() {
		if r := recover(); r != nil {
			log.Println("Could not connect", r)
		}
	}()

	flag.Parse()

	err := envconfig.Process("", conf)
	if err != nil {
		panic(err.Error())
	}

	format := "Debug: %v\nPort: %d\nTimeout: %d\n"
	log.Printf(format, conf.Debug, conf.Port, *timeout)

	err = ValidateConfiguration(*conf)
	if err != nil {
		_ = conf.CloseConnection()
		panic(err.Error())
	}
	conf.connected = true
	log.Println("Connected")

	return nil
}

func (conf *Configuration) CloseConnection() error {
	conf.connected = false

	log.Println("Connection is closed")

	return nil
}

func ValidateConfiguration(conf Configuration) error {
	err := validDBURL(conf.DBURL)

	if err != nil {
		return err
	}

	return nil
}

func validDBURL(dbUrl string) error {
	var valid = regexp.MustCompile(`postgres:\/\/(?:[^:]+):(?:[^:]+):(?:\d{4})`)

	if ok := valid.MatchString(dbUrl); ok != true {
		return errors.New("DB_URL is not set right")
	}

	return nil
}
