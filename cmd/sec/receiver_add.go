package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/fd/sec-utils/pkg/box"
	"gopkg.in/alecthomas/kingpin.v2"
)

func receiverAdd(app *kingpin.Application, configFile, name string) {
	data, err := ioutil.ReadFile(configFile)
	app.FatalIfError(err, "sec")

	r, err := box.GenerateKey()
	app.FatalIfError(err, "sec")

	rPub, err := box.PublicKey(r)
	app.FatalIfError(err, "sec")

	var conf Config
	err = json.Unmarshal(data, &conf)
	app.FatalIfError(err, "sec")

	mPub, err := box.PublicKey(conf.MasterKey)
	app.FatalIfError(err, "sec")

	if conf.ReceiverKeys == nil {
		conf.ReceiverKeys = make(map[string]string)
	}
	if _, f := conf.ReceiverKeys[name]; f {
		err := fmt.Errorf("receiver %q already exists", name)
		app.FatalIfError(err, "sec")
	}

	conf.ReceiverKeys[name] = rPub

	data, err = json.MarshalIndent(&conf, "", "  ")
	app.FatalIfError(err, "sec")

	err = ioutil.WriteFile(configFile, data, 0600)
	app.FatalIfError(err, "sec")

	fmt.Printf("%s:%s\n", r, mPub)
}
