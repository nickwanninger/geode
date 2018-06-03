package main

import (
	"flag"
	"fmt"
	// "github.com/davecgh/go-spew/spew"
	"github.com/nickwanninger/act/pkg/parser"
	"io/ioutil"
	"os"
	"strings"
)

func usage() {
	fmt.Println("Usage: act [options] <file>")
}

// if the filename passed in is a folder, look in that folder for a main.act
// if the filename is not, look for a file matching that filename, but with a .act extension
func resolveFileName(filename string) (string, error) {
	// Grab the stats of the file
	stats, err := os.Stat(filename)

	// If there was an error (file doesnt exist)
	if err != nil {
		// Try resolving the filename with .act extension
		if !strings.HasSuffix(filename, ".act") {
			return resolveFileName(filename + ".act")
		}
		// There was no file by that name, so we fail
		return "", fmt.Errorf("fatal error: No such file or directory %s", filename)
	}
	if stats.IsDir() {
		return resolveFileName(filename + "/main.act")
	}

	return filename, nil
}

func main() {
	// set the flag's usage function to my own.
	flag.Usage = usage

	flag.Parse()
	args := flag.Args()

	filename, ferr := resolveFileName(args[0])
	if ferr != nil {
		fmt.Println(ferr)
		return
	}

	data, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println(err)
	}

	lexer := parser.NewLexer()

	tokens, lexerr := lexer.Lex(data)
	if lexerr != nil {
		fmt.Println(err)
	}

	for _, t := range tokens {
		fmt.Printf("%13s - %q\n", parser.GetTokenName(t.Type), t.Value)
	}
}
