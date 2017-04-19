package main

import (
	"fmt"
	"os"
	"strings"
)

func Echo1(args []string) {
	var s, sep string
	for i := 1; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	fmt.Println(s)
}

func Echo2(args []string) {
	s, sep := "", ""
	for _, arg := range args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func Echo3(args []string) {
	fmt.Println(strings.Join(args[1:], " "))
}

func main() {
	Echo1(os.Args)
	Echo2(os.Args)
	Echo3(os.Args)
}
