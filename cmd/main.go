package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	lzstring "github.com/Lazarus/lz-string-go"
	"golang.design/x/clipboard"
)

// on frontend use https://www.npmjs.com/package/lz-string

func main() {
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}

	stat, _ := os.Stdin.Stat()
	filePath := flag.String("f", "", "file path to read from")
	flag.Parse()

	var testdata string
	if (stat.Mode()&os.ModeCharDevice) == 0 && len(os.Args) > 1 && os.Args[1] == "-" {
		var buf []byte
		scanner := bufio.NewScanner(os.Stdin)

		for scanner.Scan() {
			buf = append(buf, scanner.Bytes()...)
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		testdata = string(buf)
	} else if *filePath != "" {
		content, err := os.ReadFile(*filePath)
		if err != nil {
			log.Fatal(fmt.Errorf("error reading file: %v", err))
		}

		testdata = string(content)
	} else {
		log.Fatal("no data to compress")
	}

	keyStrUriSafe := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+-$"
	encodedData := lzstring.Compress(testdata, keyStrUriSafe)

	decodedData, err := lzstring.Decompress(encodedData, keyStrUriSafe)
	if err != nil {
		log.Fatal(fmt.Errorf("cannot ensure data valid: fail on decoding: %v", err))
	}

	if !json.Valid([]byte(decodedData)) {
		log.Fatal(fmt.Errorf("cannot ensure data valid: invalid json"))
	}

	file, err := os.CreateTemp("", "*.lzw")
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(file.Name(), []byte(encodedData), 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Encoded data: %s\n", file.Name())

	clipboard.Write(clipboard.FmtText, []byte(encodedData))
	fmt.Printf("Copied to clipboard\n")
}
