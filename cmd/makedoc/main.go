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

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "usage: makedoc [SetID] [ChallengeID]\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  Generate set[SetID]/challenge[ChallengeID]/doc.go from set[SetID]/challenge[ChallengeID]/doc.txt.")
		flag.PrintDefaults()
	}

	flag.Parse()
	if flag.NArg() != 2 {
		flag.Usage()
		os.Exit(1)
	}
	set := flag.Args()[0]
	challenge := flag.Args()[1]
	if err := run(set, challenge); err != nil {
		log.Fatal(err)
	}
}

func run(set, challenge string) (err error) {
	defer derrors.Wrap(&err, "run(%q, %q)", set, challenge)
	dir := filepath.Join(fmt.Sprintf("set%s", set), fmt.Sprintf("challenge%s", challenge))
	lines, err := readFileLines(filepath.Join(dir, "doc.txt"))
	if err != nil {
		return err
	}
	makeDoc(filepath.Join(dir, "doc.go"), set, challenge, lines)
	return nil
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

func makeDoc(filename, set, challenge string, lines []string) (err error) {
	defer derrors.Wrap(&err, "writeFile(%q, lines)", filename)
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
