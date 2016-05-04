package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/fd/sec-utils/pkg/box"
	"gopkg.in/alecthomas/kingpin.v2"
)

func doInit(app *kingpin.Application, configFile string) {
	if _, err := os.Stat(configFile); !os.IsNotExist(err) {
		err := fmt.Errorf("config file already exists")
		app.FatalIfError(err, "sec")
	}

	m, err := box.GenerateKey()
	app.FatalIfError(err, "sec")

	conf := &Config{
		MasterKey:    m,
		ReceiverKeys: map[string]string{},
	}

	data, err := json.MarshalIndent(&conf, "", "  ")
	app.FatalIfError(err, "sec")

	err = ioutil.WriteFile(configFile, data, 0600)
	app.FatalIfError(err, "sec")
}

type Config struct {
	MasterKey    string
	ReceiverKeys map[string]string
}
