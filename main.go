package main

import (
	"github.com/integrii/flaggy"

	"github.com/dontpullthis/gowipetweet/commands/tweets/delete_using_csv"
	"github.com/dontpullthis/gowipetweet/commands/tweets/dump_to_jsonl"
	"github.com/dontpullthis/gowipetweet/system/config"
)

func init() {
	flaggy.DefaultParser.ShowHelpOnUnexpected = true
	flaggy.DefaultParser.ShowHelpWithHFlag = true
}

func main() {
	var inputFile, outputFile = "", ""

	flaggy.String(&config.ConfigPath, "c", "config", "Configuration file. See config.example.yaml for more details.")

	subcommandTweetsDumpToJsonl := dump_to_jsonl.New(&inputFile, &outputFile)
	flaggy.AttachSubcommand(subcommandTweetsDumpToJsonl, 1)

	subcommandTweetsDeleteUsingCsv := delete_using_csv.New(&inputFile)
	flaggy.AttachSubcommand(subcommandTweetsDeleteUsingCsv, 2)

	flaggy.Parse()

	// config.MustInitialize must be initialized AFTER flaggy.Parse() call to receive a valid path to config file
	config.MustInitialize()

	if subcommandTweetsDumpToJsonl.Used {
		dump_to_jsonl.MustRun(inputFile, outputFile)
	} else if subcommandTweetsDeleteUsingCsv.Used {
		delete_using_csv.MustRun(inputFile)
	}
}
