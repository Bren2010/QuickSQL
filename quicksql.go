package main

import (
	"fmt"
	"os"
	"net"
	"bufio"
	"json"
	"github.com/Philio/GoMySQL"
)

// Start data declaration
var (
	// Configuration Options
	// Connection Info
	file = "/tmp/dbsock"
	workerThreads = 50
	
	// DB Connection Info
	dbHost = "localhost"
	dbUser = "user"
	dbPass = "pass"
	dbName = "example_db"
	
	// Cache Settings
	cacheAge int64 = 30 // How long cache entries should live
	cachePeriod = 1 // How often to clean the cache
	
	// System Variables
	data = make([]byte, 4096)
	listener *net.UnixListener
	db *mysql.Client
	addr *net.UnixAddr
	packetDelim = byte(4)
	unitSep = byte(31)
	connections = make(chan *net.UnixConn, 0)
)
// End

// Start GoRoutines
func main() {
	fmt.Println("Resolving address...")
	addr, err := net.ResolveUnixAddr("unix", file)
	handleErr(err, true)
	
	fmt.Println("Connecting to socket...")
	listener, err = net.ListenUnix("unix", addr)
	handleErr(err, false)
	
	if err != nil && err.String() == "listen unix "+file+": address already in use" {
		fmt.Println("Unclean shutdown!\nRemoving previous socket file...")
		err = os.Remove(file)
		handleErr(err, true)
		
		fmt.Println("Reconnecting to socket...")
		listener, err = net.ListenUnix("unix", addr)
		handleErr(err, true)
	}
	
	fmt.Println("Spawning worker threads...")
	
	i := 1
	
	for i < workerThreads {
		go workerThread()
		i = i + 1
	}
	
	fmt.Println("Done.")
	
	go sqlThread()
	go cacheCleaner()
	
	// Put main to work as a worker thread too.
	workerThread()
}

func workerThread() {
	for {
		conn, err := listener.AcceptUnix()	
		handleErr(err, true)
		
		reader := bufio.NewReader(conn)
		
		for {
			queryBytes, err := reader.ReadSlice(unitSep)
			
			if err != nil {
				break
			}
			
			cacheByte, err := reader.ReadSlice(packetDelim)
			
			if err != nil {
				break
			}
			
			received := string(queryBytes[0:len(queryBytes) - 1])
			shouldCache := string(cacheByte[0:len(cacheByte) - 1])
			
			hash := hash(received)
			
			if inArray(hash) == true && shouldCache != "1" {
				conn.Write(hashTable[hash].Data)
			} else {
				structResponse := handleQuery(received)
				
				response, err := json.Marshal(structResponse)
				handleErr(err, true)

				response = append(response, packetDelim)
				
				_, err = conn.Write(response)
				
				addToCache(hash, response)
			}
		}
	}
}
// End GoRoutines
