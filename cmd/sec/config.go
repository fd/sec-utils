package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"text/tabwriter"

	"gopkg.in/alecthomas/kingpin.v2"
)

func config(app *kingpin.Application, configFile string) {
	data, err := ioutil.ReadFile(configFile)
	app.FatalIfError(err, "sec")

	var conf Config
	err = json.Unmarshal(data, &conf)
	app.FatalIfError(err, "sec")

	w := tabwriter.NewWriter(os.Stdout, 8, 8, 2, ' ', 0)
	defer w.Flush()

	fmt.Fprintf(w, "%s\t%s\n", "master", conf.MasterKey)
	w.Flush()

	for name, key := range conf.ReceiverKeys {
		fmt.Fprintf(w, "%s\t%s\t%s\n", "receiver", name, key)
	}
}
