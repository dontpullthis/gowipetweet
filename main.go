package main

import (
	"log"

	"github.com/integrii/flaggy"

	likes_delete_using_api_list "github.com/dontpullthis/gowipetweet/commands/likes/delete_using_api_list"
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

	subcommandLikesDeleteUsingApiList := likes_delete_using_api_list.New()
	flaggy.AttachSubcommand(subcommandLikesDeleteUsingApiList, 1)

	subcommandTweetsDumpToJsonl := dump_to_jsonl.New(&inputFile, &outputFile)
	flaggy.AttachSubcommand(subcommandTweetsDumpToJsonl, 2)

	subcommandTweetsDeleteUsingCsv := delete_using_csv.New(&inputFile)
	flaggy.AttachSubcommand(subcommandTweetsDeleteUsingCsv, 3)

	subcommandTweetsToDeleteListFromJsonl := to_delete_list_from_jsonl.New(&inputFile, &outputFile, &expression)
	flaggy.AttachSubcommand(subcommandTweetsToDeleteListFromJsonl, 4)

	flaggy.Parse()

	if subcommandLikesDeleteUsingApiList.Used {
		likes_delete_using_api_list.MustRun()
	} else if subcommandTweetsDumpToJsonl.Used {
		dump_to_jsonl.MustRun(inputFile, outputFile)
	} else if subcommandTweetsDeleteUsingCsv.Used {
		delete_using_csv.MustRun(inputFile)
	} else if subcommandTweetsToDeleteListFromJsonl.Used {
		to_delete_list_from_jsonl.MustRun(inputFile, outputFile, expression)
	}

	log.Println("Done.")
}
