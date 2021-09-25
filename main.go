package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/julieqiu/derrors"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "usage: cryptopals [SetID] [ChallengeID]\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  From https://cryptopals.com/sets/<SetID>/challenges/<ChallengeID>\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  For example, to run https://cryptopals.com/sets/1/challenges/7:\n")
		fmt.Fprintf(flag.CommandLine.Output(), "	cryptopals 1 7\n")
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
		return err
	}
}

func run(set, challenge string) {
	switch set {
	case 1:
		switch challenge {
		case 1:
			return setc1()
		default:
			fmt.Println("Still working on it!")
		}
	default:
		fmt.Println("Still working on it!")
	}
}

func setc1() (err error) {
	defer derrors.Wrap(&err, "setc1: error running solution for Set1 Challenge1")
	return nil
}
