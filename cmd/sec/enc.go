package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"

	"github.com/fd/sec-utils/pkg/box"
	"gopkg.in/alecthomas/kingpin.v2"
)

func enc(app *kingpin.Application, configFile string) {
	data, err := ioutil.ReadFile(configFile)
	app.FatalIfError(err, "sec")

	var conf Config
	err = json.Unmarshal(data, &conf)
	app.FatalIfError(err, "sec")

	var rKeys []string
	for _, k := range conf.ReceiverKeys {
		rKeys = append(rKeys, k)
	}

	data, err = ioutil.ReadAll(os.Stdin)
	app.FatalIfError(err, "sec")

	encdata, err := box.Seal(data, conf.MasterKey, rKeys)
	app.FatalIfError(err, "sec")

	_, err = io.Copy(os.Stdout, bytes.NewReader(encdata))
	app.FatalIfError(err, "sec")
}
