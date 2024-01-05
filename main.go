package main

import "fmt"

var Version string
var BuildTime string

func main() {
	fmt.Printf("Version: %s\n", Version)
	fmt.Printf("Buildtime: %s\n", BuildTime)
}
