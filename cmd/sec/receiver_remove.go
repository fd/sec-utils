package main

import (
	"encoding/json"
	"io/ioutil"

	"gopkg.in/alecthomas/kingpin.v2"
)

func receiverRemove(app *kingpin.Application, configFile, name string) {
	data, err := ioutil.ReadFile(configFile)
	app.FatalIfError(err, "sec")

	var conf Config
	err = json.Unmarshal(data, &conf)
	app.FatalIfError(err, "sec")

	if conf.ReceiverKeys == nil {
		conf.ReceiverKeys = make(map[string]string)
	}

	delete(conf.ReceiverKeys, name)

	data, err = json.MarshalIndent(&conf, "", "  ")
	app.FatalIfError(err, "sec")

	err = ioutil.WriteFile(configFile, data, 0600)
	app.FatalIfError(err, "sec")
}
