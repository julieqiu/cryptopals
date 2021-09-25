package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/julieqiu/derrors"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "usage: makedoc [dir]\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  Convert [dir]/doc.txt to [dir]/doc.go.")
		flag.PrintDefaults()
	}

	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}
	dir := flag.Args()[0]
	if err := run(dir); err != nil {
		log.Fatal(err)
	}
}

func run(dir string) (err error) {
	defer derrors.Wrap(&err, "run(%q)", dir)
	set, challenge, err := setAndChallengeFromDir(dir)
	if err != nil {
		return err
	}
	lines, err := readFileLines(filepath.Join(dir, "doc.txt"))
	if err != nil {
		return err
	}
	makeDoc(filepath.Join(dir, "doc.go"), set, challenge, lines)
	return nil
}

func setAndChallengeFromDir(dir string) (set string, challenge string, err error) {
	defer derrors.Wrap(&err, "setAndChallengeFromDir(%q)", dir)
	parts := strings.Split(dir, "s")
	if len(parts) != 2 {
		return "", "", fmt.Errorf("invalid directory, no set")
	}
	parts = strings.Split(parts[1], "c")
	if len(parts) != 2 {
		return "", "", fmt.Errorf("invalid directory, no challenge")
	}
	return parts[0], parts[1], nil
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
//     go run ./cmd/makedoc s%[1]sc%[2]s
//

`, set, challenge)
	for _, l := range lines {
		content += fmt.Sprintf("// %s\n", l)
	}
	content += "//\n"
	content += "//\n"
	content += fmt.Sprintf("package s%[1]sc%[2]s", set, challenge)
	return ioutil.WriteFile(filename, []byte(content), 0644)
}
