//C01E01 prints its command line arguments one per line along with their index 

package main

import (
	"fmt"
	"os"
)

func main() {
	for i := 0; i < len(os.Args); i++ {
		fmt.Println(i, os.Args[i])
	}
}
