//C01E01 prints its command line arguments and the name of the program

package main

import (
	"fmt"
	"os"
)

func main() {

	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
