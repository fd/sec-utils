package main

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/fd/sec-utils/pkg/box"
	"gopkg.in/alecthomas/kingpin.v2"
)

func dec(app *kingpin.Application, creds string) {
	keys := strings.SplitN(creds, ":", 2)
	if len(keys) != 2 {
		err := errors.New("invalid credentials")
		app.FatalIfError(err, "sec")
	}

	data, err := ioutil.ReadAll(os.Stdin)
	app.FatalIfError(err, "sec")

	decdata, err := box.Open(data, keys[0], keys[1])
	app.FatalIfError(err, "sec")

	_, err = io.Copy(os.Stdout, bytes.NewReader(decdata))
	app.FatalIfError(err, "sec")
}
