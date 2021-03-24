package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/alexflint/go-arg"
	ppt "github.com/zerodoctor/goprettyprinter"
)

var args struct {
	Old        string `arg:"required,-o,--old"`
	New        string `arg:"required,-n,--new"`
	DateFormat string `arg:"-d,--date"`
	StartLine  int    `arg:"-s,--start" default:"0"`
	EndLine    int    `arg:"-e,--end" default:"1"`
}

func getStdin() (string, error) {
	info, err := os.Stdin.Stat()
	if err != nil {
		log.Fatalln(err.Error())
	}

	if info.Mode()&os.ModeNamedPipe == 0 {
		ppt.Errorln()
		return "", errors.New("must be used with pipe")
	}

	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return "", errors.New("failed to read input")
	}

	return string(input), nil
}

func main() {
	arg.MustParse(&args)
	input, err := getStdin()
	if err != nil {
		ppt.Errorln(err.Error())
	}

	lineSplit := strings.Split(input, "\n")

	for i := args.StartLine; i < args.EndLine; i++ {
		new := args.New
		if new == "\\b" {
			new = ""
		}
		lineSplit[i] = strings.ReplaceAll(lineSplit[i], args.Old, new)
	}

	fmt.Println(strings.Join(lineSplit, "\n"))
}
