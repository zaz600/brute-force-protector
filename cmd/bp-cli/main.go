package main

import (
	"os"
)

type ListType string

const (
	Black ListType = "BlackList"
	White ListType = "WhiteList"
)

func main() {
	os.Exit(CLI(os.Args))
}
