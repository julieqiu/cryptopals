package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/julieqiu/derrors"
)

var mode = flag.String("m", "doc",
	`set: Generate templates for set[SetID]/challenge[ChallengeID]
doc: Generate set[SetID]/challenge[ChallengeID]/doc.go from set[SetID]/challenge[ChallengeID]/doc.txt`)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "usage: makedoc [SetID] [ChallengeID]\n")
		flag.PrintDefaults()
	}

	flag.Parse()
	if flag.NArg() != 2 {
		flag.Usage()
		os.Exit(1)
	}
	set := flag.Args()[0]
	challenge := flag.Args()[1]
	mode := *mode
	if err := run(set, challenge, mode); err != nil {
		log.Fatal(err)
	}
}

func run(set, challenge, mode string) (err error) {
	defer derrors.Wrap(&err, "run(%q, %q)", set, challenge)
	dir := filepath.Join(fmt.Sprintf("set%s", set), fmt.Sprintf("challenge%s", challenge))
	switch mode {
	case "doc":
		return makeDoc(dir, set, challenge)
	case "set":
		if err := makeSet(dir, set, challenge); err != nil {
			return err
		}
		return makeDoc(dir, set, challenge)
	default:
		return fmt.Errorf("-m flag is invalid")
	}
}

// readfilelines reads and returns the lines from a file.
func readFileLines(filename string) (lines []string, err error) {
	defer derrors.Wrap(&err, "readFileLines(%q)", filename)
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()
		lines = append(lines, line)
	}
	if s.Err() != nil {
		return nil, s.Err()
	}
	return lines, nil
}

func makeDoc(dir, set, challenge string) (err error) {
	defer derrors.Wrap(&err, "makeDoc(%q, lines)", dir)
	lines, err := readFileLines(filepath.Join(dir, "doc.txt"))
	if err != nil {
		return err
	}

	filename := filepath.Join(dir, "doc.go")
	content := fmt.Sprintf(`// This file is generated using cmd/makedoc. DO NOT EDIT.
// To update, edit the doc.txt file in this directory.
// Then run
//     go run ./cmd/makedoc %s %s
//

`, set, challenge)
	for _, l := range lines {
		content += fmt.Sprintf("// %s\n", l)
	}
	content += "//\n"
	content += "//\n"
	content += fmt.Sprintf("package challenge%s", challenge)
	return ioutil.WriteFile(filename, []byte(content), 0644)
}

func makeSet(dir, set, challenge string) (err error) {
	defer derrors.Wrap(&err, "makeSet(%q, %q, %q)", dir, set, challenge)

	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return fmt.Errorf("os.MkdirAll(%q, %d): %v", dir, os.ModePerm, err)
	}

	if err := makeTemplateDoc(dir, set, challenge); err != nil {
		return err
	}
	return makeTemplateExample(dir, set, challenge)
}

func makeTemplateDoc(dir, set, challenge string) (err error) {
	defer derrors.Wrap(&err, "makeTemplateDoc(%q, %q, %q)", dir, set, challenge)

	filename := filepath.Join(dir, "doc.txt")
	content := fmt.Sprintf(`This package provides a solution to https://cryptopals.com/sets/%[1]s/challenges/%[2]s.

Problem

Coming soon!

Solution

Coming soon!`, set, challenge)
	return ioutil.WriteFile(filename, []byte(content), 0644)
}

func makeTemplateExample(dir, set, challenge string) (err error) {
	defer derrors.Wrap(&err, "makeTemplateExample(%q, %q, %q)", dir, set, challenge)

	filename := filepath.Join(dir, "example_test.go")
	content := fmt.Sprintf(`package challenge%s_test

func Example() {
}`, challenge)
	return ioutil.WriteFile(filename, []byte(content), 0644)
}
