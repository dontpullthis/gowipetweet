package main

import (
	"github.com/integrii/flaggy"

	"github.com/dontpullthis/gowipetweet/commands/tweets/delete_using_csv"
	"github.com/dontpullthis/gowipetweet/system/config"
)

func init() {
	flaggy.DefaultParser.ShowHelpOnUnexpected = true
	flaggy.DefaultParser.ShowHelpWithHFlag = true
}

func main() {
	var inputFile = ""

	flaggy.String(&config.ConfigPath, "c", "config", "Configuration file. See config.example.yaml for more details.")

	subcommandTweetsDeleteUsingCsv := delete_using_csv.New(&inputFile)
	flaggy.AttachSubcommand(subcommandTweetsDeleteUsingCsv, 1)

	flaggy.Parse()

	// config.MustInitialize must be initialized AFTER flaggy.Parse() call to receive a valid path to config file
	config.MustInitialize()

	if subcommandTweetsDeleteUsingCsv.Used {
		delete_using_csv.MustRun(inputFile)
	}
}
