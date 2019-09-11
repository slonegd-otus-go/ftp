package main

import (
	"fmt"
	"log"
	"time"

	"github.com/jlaffaye/ftp"
)

func main() {
	connection, err := ftp.Dial("localhost:21", ftp.DialWithTimeout(1*time.Second))
	if err != nil {
		log.Fatalf("dial failed: %s", err)
	}
	defer func() {
		if err = connection.Quit(); err != nil {
			log.Fatalf("connection quit failed: %s", err)
		}
	}()

	err = connection.Login("username", "")
	if err != nil {
		log.Fatalf("connection login failed: %s", err)
	}

	list, err := connection.List("/")
	if err != nil {
		log.Fatalf("connection list failed: %s", err)
	}

	fmt.Printf("%+s", list)
}
