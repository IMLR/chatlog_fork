package main

import (
	"log"

	"github.com/IMLR/chatlog_fork/cmd/chatlog"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	chatlog.Execute()
}
