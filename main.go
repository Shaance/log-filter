package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	logLevels := map[string]bool{"INFO": true, "DEBUG": true, "ERROR": true, "CRITICAL": true}
	path, logLevel := getProgramArguments()

	if !logLevels[strings.ToUpper(logLevel)] {
		log.Fatal(logLevel + " value not recognized. See the -level option usage to learn more about available options")
	}

	file, err := os.Open(path)
	defer file.Close()
	terminateOnError(err)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, strings.ToUpper(logLevel)) {
			fmt.Println(line)
		}
	}

	terminateOnError(scanner.Err())

}

func getProgramArguments() (string, string) {
	path := flag.String("path", "logfile.log", "The path to the log file to be filtered.")
	logLevel := flag.String("level", "ERROR", "Log level to search for. "+
		"Options are DEBUG, INFO, ERROR and CRITICAL")
	flag.Parse()

	return *path, *logLevel
}

func terminateOnError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
