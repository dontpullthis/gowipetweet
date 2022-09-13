package delete_using_csv

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/dontpullthis/gowipetweet/client/twitter"
	"github.com/dontpullthis/gowipetweet/system/config"
)

func MustRun(inputFile string) {
	f, err := os.Open(inputFile)
	if err != nil {
		log.Fatal("Unable to read input file "+inputFile+". ", err)
	}
	defer f.Close()
	csvReader := csv.NewReader(f)

	cfg := config.GetConfig()
	if twitter.MustInitialize(cfg) {
		config.MustSave()
	}

	for {
		var tweetIdLine, err = csvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal("Failed to read a line from file "+inputFile+". ", err)
		}

		fmt.Printf("Deleting tweet %s...\n", tweetIdLine[0])
		twitter.ClientInstance.MustDeleteTweet(tweetIdLine[0])
	}
}
