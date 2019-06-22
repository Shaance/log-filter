# Log filter

## Description 
Simple program written in Go to filter log file on log level basis - made for training purposes.

## Requirements
You just need to have [Go](https://golang.org/dl/) installed.

## How to run
Go to the project root directory and launch `go run .` command. You can also pass some options.
```
  -level string
        Log level to search for. Options are DEBUG, INFO, ERROR and CRITICAL (default "ERROR")
  -path string
        The path to the log file to be filtered. (default "logfile.log")

```
