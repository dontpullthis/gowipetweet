package to_delete_list_from_jsonl

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
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
	for scanner.Scan() {
		line := scanner.Text()

		var record Record
		err := json.Unmarshal([]byte(line), &record)
		if err != nil {
			log.Fatal("Unable to read marshal a line: "+line+". ", err)
		}
		fmt.Println(record.Tweet.Id + "\t" + record.Tweet.FullText)
	}
	// writer := bufio.NewWriter(out)

	// err = converter.JavascriptToJSONL(scanner, writer)
	// if err != nil {
	// 	log.Fatal("Failed to convert the '"+inputFile+"' to JSON Lines. ", err)
	// }
}
