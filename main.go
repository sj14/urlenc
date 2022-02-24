package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
)

func main() {
	decodeFlag := flag.Bool("decode", false, "perform decoding instead of encoding of the given input")
	flag.Parse()

	input, err := readInput()
	if err != nil {
		log.Fatalln(err)
	}

	result := ""

	if *decodeFlag {
		result, err = url.QueryUnescape(input)
		if err != nil {
			log.Fatalln(err)
		}
	} else {
		result = url.QueryEscape(input)
	}

	fmt.Println(result)
}

// read program input from stdin or argument
func readInput() (string, error) {
	// from stdin/pipe
	if flag.NArg() == 0 {

		// check if it's piped or from empty stdin
		// https://stackoverflow.com/a/26567513
		stat, err := os.Stdin.Stat()
		if err != nil {
			return "", fmt.Errorf("failed to get stdin stats: %v", err)
		}
		if (stat.Mode() & os.ModeCharDevice) != 0 {
			return "", nil
		}

		// read the input from the pipe
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			return "", fmt.Errorf("failed to read input: %v", err)
		}
		return strings.TrimSpace(input), nil
	}

	// from argument
	if flag.NArg() > 1 {
		return "", fmt.Errorf("takes at most one input")
	}

	return flag.Arg(0), nil
}
