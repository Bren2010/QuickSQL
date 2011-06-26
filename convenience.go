package main

// Start data declaration
import (
	"fmt"
	"os"
)
// End

// Start functions
// A quick command so I don't have to have if statements after every line.
func handleErr(err os.Error, critical bool) {
	if err != nil {
		fmt.Println("Error: ", err)
		
		if critical == true {
			os.Exit(-1)
		}
	}
}
// End functions
