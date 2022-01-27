package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 0 {
		printArgs(args)
		printFlag()
	} else {
		println("hello world")
	}
}
/**
	运行：
		go build main.go
		./main
 */

// 方式1
func printArgs(args []string) {
	for _,v := range args {
		fmt.Printf("args is %s \n", v)
	}
}

// 方式2
func printFlag() {
	name := flag.String("name", "", "string type")
	flag.Parse()
	fmt.Printf("flag name is %s \n", *name)
}