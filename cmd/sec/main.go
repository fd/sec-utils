package main

import (
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
	"limbo.services/version"
)

func main() {
	var (
		configFile   string
		receiverName string
		receiverKey  string
	)

	app := kingpin.New("sec", "Simple secure data").
		Version(version.Get().String()).
		Author(version.Get().ReleasedBy)

	initCmd := app.Command("init", "Initialize a config file")
	initCmd.Flag("config", "path to the config file").Short('c').
		Default("./sec-config.json").PlaceHolder("FILE").Envar("SEC_CONFIG").
		StringVar(&configFile)

	configCmd := app.Command("config", "Show the content of a config file")
	configCmd.Flag("config", "path to the config file").Short('c').
		Default("./sec-config.json").PlaceHolder("FILE").Envar("SEC_CONFIG").
		ExistingFileVar(&configFile)

	receiverAddCmd := app.Command("receiver-add", "Register a new receiver")
	receiverAddCmd.Arg("name", "name of the receiver").Required().
		StringVar(&receiverName)
	receiverAddCmd.Flag("config", "path to the config file").Short('c').
		Default("./sec-config.json").PlaceHolder("FILE").Envar("SEC_CONFIG").
		ExistingFileVar(&configFile)

	receiverRemoveCmd := app.Command("receiver-remove", "Remove a receiver")
	receiverRemoveCmd.Arg("name", "name of the receiver").Required().
		StringVar(&receiverName)
	receiverRemoveCmd.Flag("config", "path to the config file").Short('c').
		Default("./sec-config.json").PlaceHolder("FILE").Envar("SEC_CONFIG").
		ExistingFileVar(&configFile)

	encCmd := app.Command("enc", "Encrypt data from STDIN")
	encCmd.Flag("config", "path to the config file").Short('c').
		Default("./sec-config.json").PlaceHolder("FILE").Envar("SEC_CONFIG").
		ExistingFileVar(&configFile)

	decCmd := app.Command("dec", "Decrypt data from STDIN")
	decCmd.Flag("key", "receiver key").Short('k').
		PlaceHolder("KEY").Envar("SEC_KEY").
		Required().StringVar(&receiverKey)

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case initCmd.FullCommand():
		doInit(app, configFile)
	case configCmd.FullCommand():
		config(app, configFile)
	case receiverAddCmd.FullCommand():
		receiverAdd(app, configFile, receiverName)
	case receiverRemoveCmd.FullCommand():
		receiverRemove(app, configFile, receiverName)
	case encCmd.FullCommand():
		enc(app, configFile)
	case decCmd.FullCommand():
		dec(app, receiverKey)
	}
}
