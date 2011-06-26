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

// A function to always make sure the connection is closed.
func close() {
	err := listener.Close()
	handleErr(err, false)
	fmt.Println("Listener Closed.")
}
// End functions
