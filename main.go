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

	subcommandTweetsDeleteUsingCsv := flaggy.NewSubcommand("tweets:delete:using_csv")
	subcommandTweetsDeleteUsingCsv.Description = "Deletes tweets using a CSV file as a data source"
	subcommandTweetsDeleteUsingCsv.String(&inputFile, "i", "input-file", "Path to CSV file where each line is ID of tweet to delete")
	flaggy.AttachSubcommand(subcommandTweetsDeleteUsingCsv, 1)

	flaggy.Parse()

	config.MustInitialize()

	if subcommandTweetsDeleteUsingCsv.Used {
		delete_using_csv.MustRun(inputFile)
	}
}
