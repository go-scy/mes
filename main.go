package main

import "fmt"

// build vars
var (
	version, commitHash string
)

func main() {
	fmt.Println("version:", version)
	fmt.Println("commit:", commitHash)
}
