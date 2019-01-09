package main

import (
	"flag"
	"os"
)

type CLI struct {
	bc *Blockchain
}

// Run parses command line arguments and processes commands
func (cli *CLI) Run() {
	cli.validateArgs()

	// using the standard flag package to parse command-line arguments.
	addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError) //blockchain_go addblock "Pay 0.031337 for a coffee"
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError) //blockchain_go printchain
	addBlockData := addBlockCmd.String("data", "", "Block data")

	switch os.Args[1] {
	case "addblock":
		err := addBlockCmd.Parse(os.Args[2:])
	case "printchain":
		err := printChainCmd.Parse(os.Args[2:])
	default:
		cli.printUsage()
		os.Exit(1)
	}

	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			addBlockCmd.Usage()
			os.Exit(1)
		}
		cli.addBlock(*addBlockData)
	}

	if printChainCmd.Parsed() {
		cli.printChain()
	}
}