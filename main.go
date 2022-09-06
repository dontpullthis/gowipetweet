package main

import (
	"github.com/integrii/flaggy"

	"github.com/dontpullthis/gowipetweet/commands/tweets/delete_using_csv"
)

func main() {
	flaggy.DefaultParser.ShowHelpOnUnexpected = false
	flaggy.DefaultParser.ShowHelpWithHFlag = true
	// Declare variables and their defaults
	var inputFile = ""

	// Create the subcommand
	subcommandTweetsDeleteUsingCsv := flaggy.NewSubcommand("tweets:delete:using_csv")
	subcommandTweetsDeleteUsingCsv.Description = "Deletes tweets using a CSV file as a data source"
	subcommandTweetsDeleteUsingCsv.String(&inputFile, "c", "csv-file", "Path to CSV file where each line is ID of tweet to delete")
	flaggy.AttachSubcommand(subcommandTweetsDeleteUsingCsv, 1)

	// Parse the subcommand and all flags
	flaggy.Parse()

	if subcommandTweetsDeleteUsingCsv.Used {
		delete_using_csv.Run(inputFile)
	}
}
