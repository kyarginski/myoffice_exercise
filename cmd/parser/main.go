package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"myoffice/internal/processes"
)

var (
	sourceFlag   string
	parallelFlag string
	helpFlag     string
)

var errUsage = errors.New("usage")

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `Usage: %s [option...] 

Options:
`,
			os.Args[0])
		flag.PrintDefaults()
	}
	flag.StringVar(&helpFlag, "help", "", "show help about arguments")
	flag.StringVar(&sourceFlag, "source", "", "name of input data files")
	flag.StringVar(&parallelFlag, "parallel", "0", "parallel processing N = count of goroutines, 0 = count of CPU cores")
	flag.Parse()
	if len(helpFlag) > 0 {
		flag.Usage()
		os.Exit(2)
	}

	err := run()
	if err != nil {
		if errors.Is(err, errUsage) {
			flag.Usage()
			os.Exit(2)
		}
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}

func run() error {
	if len(sourceFlag) == 0 {
		return errUsage
	}
	return processes.Run(sourceFlag, parallelFlag)
}
