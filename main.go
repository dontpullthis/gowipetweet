package main

import (
	"github.com/integrii/flaggy"

	"github.com/dontpullthis/gowipetweet/client/twitter"
	"github.com/dontpullthis/gowipetweet/commands/tweets/delete_using_csv"
	"github.com/dontpullthis/gowipetweet/system/config"
)

func main() {
	flaggy.DefaultParser.ShowHelpOnUnexpected = false
	flaggy.DefaultParser.ShowHelpWithHFlag = true

	var configFile = "config.yaml"
	var inputFile = ""

	flaggy.String(&configFile, "c", "config", "Configuration file. See config.example.yaml for more details.")

	subcommandTweetsDeleteUsingCsv := flaggy.NewSubcommand("tweets:delete:using_csv")
	subcommandTweetsDeleteUsingCsv.Description = "Deletes tweets using a CSV file as a data source"
	subcommandTweetsDeleteUsingCsv.String(&inputFile, "i", "input-file", "Path to CSV file where each line is ID of tweet to delete")
	flaggy.AttachSubcommand(subcommandTweetsDeleteUsingCsv, 1)

	flaggy.Parse()

	cfg := config.MustInitialize(configFile)

	if twitter.MustInitialize(&cfg) {
		config.MustSave(configFile, cfg)
	}

	if subcommandTweetsDeleteUsingCsv.Used {
		delete_using_csv.MustRun(inputFile)
	}
}
