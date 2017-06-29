package main

import (
	"log"
	"os"
)

func main() {
	os.Exit(Main())
}

func Main() int {
	term := make(chan os.Signal)

	webServer := NewWebServer()
	go webServer.Run()

	tailer := NewTailer()
	go tailer.Run()

	parser := NewParser(tailer.Out)
	go parser.Run()

	select {
	case <-term:
		log.Println("Received SIGTERM, exiting gracefully...")
	case err := <-webServer.ListenError():
		log.Println("Error starting web server, exiting gracefully:", err)
	case err := <-tailer.ListenError():
		log.Println("Error starting tail, exiting gracefully:", err)
	case err := <-parser.ListenError():
		log.Println("Error starting parser, exiting gracefully:", err)
	}
	return 0
}
