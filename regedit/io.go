package regedit

import (
	"bufio"
	"fmt"
	"golang.org/x/text/encoding"
	"io"
	"log"
	"os/exec"
)

func errorLog(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func closeReader(reader io.ReadCloser) {
	err := reader.Close()
	if err != nil {
		errorLog(err)
	}
}

// path : window-tool path
// si : start index
// decoder : korean.EUCKR.NewDecoder()
func command(data []string, path string, si int, decoder *encoding.Decoder) (strings []string, err error) {
	var reader io.ReadCloser
	// start command
	cmd := exec.Command("reg", "query", path)
	if reader, err = cmd.StdoutPipe(); err != nil {
		return
	}

	//// close reader
	//defer closeReader(reader)

	// start cmd
	if err = cmd.Start(); err != nil {
		return
	}

	decoder.Reset()
	scanner := bufio.NewScanner(decoder.Reader(reader))
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > si {
			data = append(data, line[si:])
		}
	}

	if err := cmd.Wait(); err != nil {
		fmt.Println("Error:", err)
		panic(err)
	}

	return data, err
}
