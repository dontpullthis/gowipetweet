package dump_to_jsonl

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func MustRun(inputFile string, outputFile string) {
	in, err := os.Open(inputFile)
	if err != nil {
		log.Fatal("Unable to read input file "+inputFile+". ", err)
	}
	defer in.Close()

	out, err := os.Create(outputFile)
	if err != nil {
		log.Fatal("Unable to read output file "+inputFile+". ", err)
	}
	defer out.Close()

	scanner := bufio.NewScanner(in)

	scanner.Scan()
	failOnScannerError(scanner)

	nestingLevel := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		nestingLevel = nestingLevel + strings.Count(line, "{") - strings.Count(line, "}")
		if 0 == nestingLevel {
			if strings.HasSuffix(line, ",") { // end if element
				line = line[:len(line)-1] + "\n"
			} else if strings.HasSuffix(line, "]") {
				// close square bracked at the end of file
				// open bracket was eliminated by first invocation of scanner.Scan() method outside of loop
				break
			}
		}
		out.Write([]byte(line))
	}
	failOnScannerError(scanner)
}

func failOnScannerError(scanner *bufio.Scanner) {
	if err := scanner.Err(); err != nil {
		log.Fatal("An error occurret while reading the input file. ", err)
	}
}
