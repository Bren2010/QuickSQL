package main

import (
	"fmt"
	"os"
	"net"
	"os/signal"
	"strings"
	"encoding/hex"
	"bufio"
	"json"
	"github.com/Philio/GoMySQL"
)

// Start data declaration
var (
	// Configuration Options
	// Connection Info
	file = "/tmp/dbsock"
	
	// DB Connection Info
	dbHost = "localhost"
	dbUser = "user"
	dbPass = "pass"
	dbName = "example_db"
	
	
	// System Variables
	data = make([]byte, 4096)
	listener *net.UnixListener
	db *mysql.Client
	addr *net.UnixAddr
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
	
	go signalHandler()
	go sqlThread()
	go cacheCleaner()
	
	for {
		conn, err := listener.AcceptUnix()	
		handleErr(err, true)
		
		reader := bufio.NewReader(conn)
		
		dst := make([]byte, 4096)
		
		for {
			line, _, err := reader.ReadLine()
			
			if string(line) != "" {
				c, err := hex.Decode(dst, line)
				handleErr(err, true)
				
				received := string(dst[0:c])
				
				hash := hash(received)
				
				if inArray(hash) == true {
					conn.Write(hashTable[hash].Data)
				} else {
					structResponse := handleQuery(received)
					
					jsonResponse, err := json.Marshal(structResponse)
					handleErr(err, true)
					fmt.Println(string(jsonResponse))
					response := hex.EncodeToString(jsonResponse)
					
					c, err = conn.Write([]byte(response+"\n\r"))
					
					addToCache(hash, response+"\n\r")
				}
			} else if err != nil {
				break
			}
		}
	}
}

func signalHandler() {
	for true {
		sig := <- signal.Incoming
		array := strings.Split(sig.String(), ": ", 2)
		
		if array[0] != "SIGCHLD" {
			fmt.Println("Received "+array[0]+".")
			close()
			os.Exit(0)
		}
	}
}
// End GoRoutines
