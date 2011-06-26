package main

import (
	"fmt"
	"os"
	"github.com/Philio/GoMySQL"
)

// Start data declaration
var (
	queries = make(chan SqlRequest)
	responses = make(chan SqlOutgoing)
)

type SqlOutgoing struct {
	Error string
	AffectedRows uint64
	LastInsertID uint64
	Count int
	Results []mysql.Map
}

type SqlRequest struct {
	Sql string
	ResponseChan chan SqlOutgoing
}
// End

// Start GoRoutines
func sqlThread() {
	// Connect to DB
	db, err := mysql.DialTCP(dbHost, dbUser, dbPass, dbName)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for true {
		db.Start()
		
		// Read queries from channel and execute until the end of time.
		queryStruct := <- queries 

		
		// Run query through MySQL
		err = db.Query(queryStruct.Sql)
		handleErr(err, false)
		
		// Get result set
		result, err := db.UseResult()
		handleErr(err, true)
		
		var rowArray []mysql.Map
		
		// Get each row from the result and perform some processing
		for {
			row := result.FetchMap()

			// Quit reading if there are no more rows.
			if row == nil {
				break
			}
			
			// Add row to result slice.
			rowArray = append(rowArray, row)
		}
		
		// Get the error as a string or a return blank string.
		errString := ""
		if err != nil {
			errString = err.String()
		}
		
		// Create struct to response.
		response := SqlOutgoing{errString, db.AffectedRows, db.LastInsertId, len(rowArray), rowArray}
		
		// Reply with struct.
		queryStruct.ResponseChan <- response
		
		// Clean up.
		db.FreeResult()
	}
}
// End GoRoutines

// Start functions
// handleQuery() forces asynchronous MySQL queries to be thread safe.
func handleQuery(received string) SqlOutgoing {
	// Create channel to receive the response.
	responseChan := make(chan SqlOutgoing, 4096)
	
	// Create request and send it.
	request := SqlRequest{received, responseChan}
	queries <- request
	
	// Wait for response from the sql thread and return it.
	response := <- responseChan
	return response
}
// End functions
