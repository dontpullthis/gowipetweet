package main

import (
	"log"

	"github.com/integrii/flaggy"

	"github.com/dontpullthis/gowipetweet/commands/tweets/delete_using_csv"
	"github.com/dontpullthis/gowipetweet/commands/tweets/dump_to_jsonl"
	"github.com/dontpullthis/gowipetweet/commands/tweets/to_delete_list_from_jsonl"
	"github.com/dontpullthis/gowipetweet/system/config"
)

func init() {
	flaggy.DefaultParser.ShowHelpOnUnexpected = true
	flaggy.DefaultParser.ShowHelpWithHFlag = true
}

func main() {
	var inputFile, outputFile, expression = "", "", ""

	log.Println("Starting...")

	flaggy.SetName("gowipetweet")
	flaggy.SetDescription("A utility for tweet deletion. Webpage: https://github.com/dontpullthis/gowipetweet")

	flaggy.String(&config.ConfigPath, "c", "config", "Configuration file. See config.example.yaml for more details.")

	subcommandTweetsDumpToJsonl := dump_to_jsonl.New(&inputFile, &outputFile)
	flaggy.AttachSubcommand(subcommandTweetsDumpToJsonl, 1)

	subcommandTweetsDeleteUsingCsv := delete_using_csv.New(&inputFile)
	flaggy.AttachSubcommand(subcommandTweetsDeleteUsingCsv, 1)

	subcommandTweetsToDeleteListFromJsonl := to_delete_list_from_jsonl.New(&inputFile, &outputFile, &expression)
	flaggy.AttachSubcommand(subcommandTweetsToDeleteListFromJsonl, 1)

	flaggy.Parse()

	if subcommandTweetsDumpToJsonl.Used {
		dump_to_jsonl.MustRun(inputFile, outputFile)
	} else if subcommandTweetsDeleteUsingCsv.Used {
		delete_using_csv.MustRun(inputFile)
	} else if subcommandTweetsToDeleteListFromJsonl.Used {
		to_delete_list_from_jsonl.MustRun(inputFile, outputFile, expression)
	}

	log.Println("Done.")
}
