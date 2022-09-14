package dump_to_jsonl

import (
	"bufio"
	"log"
	"os"

	"github.com/dontpullthis/gowipetweet/converter"
)

func MustRun(inputFile string, outputFile string) {
	in, err := os.Open(inputFile)
	if err != nil {
		log.Fatal("Unable to read input file "+inputFile+". ", err)
	}
	defer in.Close()

	out, err := os.Create(outputFile)
	if err != nil {
		log.Fatal("Unable to read output file "+outputFile+". ", err)
	}
	defer out.Close()

	scanner := bufio.NewScanner(in)
	writer := bufio.NewWriter(out)

	err = converter.JavascriptToJSONL(scanner, writer)
	if err != nil {
		log.Fatal("Failed to convert the '"+inputFile+"' to JSON Lines. ", err)
	}
}
