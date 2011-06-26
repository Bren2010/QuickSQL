package main

// Start data declaration
import (
	"fmt"
	"os"
	"github.com/Philio/GoMySQL"
)

var (
	queries = make(chan SqlRequest)
	responses = make(chan SqlOutgoing)
)

type SqlOutgoing struct {
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
		db.Lock()
		db.Start()
		
		queryStruct := <- queries // Read queries from channel and execute until the end of time.
		query := queryStruct.Sql
		
		err = db.Query(query)
		handleErr(err, true)
		
		// Get result set
		result, err := db.UseResult()
		handleErr(err, true)
		
		var rowArray []mysql.Map
		
		// Get each row from the result and perform some processing
		for {
			row := result.FetchMap()

			if row == nil {
				break
			}
			
			rowArray = append(rowArray, row)
		}
		
		response := SqlOutgoing{len(rowArray), rowArray}
		
		queryStruct.ResponseChan <- response
		
		db.FreeResult()
		db.Unlock()
	}
}
// End GoRoutines

// Start functions
func handleQuery(received string) SqlOutgoing {
	responseChan := make(chan SqlOutgoing, 4096)
	
	request := SqlRequest{received, responseChan}
	queries <- request
	
	response := <- responseChan
	return response
}
// End functions
