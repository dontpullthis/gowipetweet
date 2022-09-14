package converter

import (
	"bufio"
	"strings"
)

func JavascriptToJSONL(s *bufio.Scanner, w *bufio.Writer) error {
	s.Scan()
	if s.Err() != nil {
		return s.Err()
	}

	nestingLevel := 0
	for s.Scan() {
		line := strings.TrimSpace(s.Text())

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
		_, err := w.Write([]byte(line))
		if err != nil {
			return err
		}
	}

	if s.Err() != nil {
		return s.Err()
	}

	err := w.Flush()
	if err != nil {
		return err
	}

	return nil
}
