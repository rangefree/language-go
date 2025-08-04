package main

import (
	"log"
	"os"
)

func main() {
	log.Println("This is log message")

	log.SetPrefix("SomePrefix")
	log.Println("Prefixed log message")

	//Log Flags:
	log.SetPrefix("PREFIX: ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
	log.Println("Text of the log message.")

	info.Println("Message for info logger")
	err.Println("Message for error logger")

	// log to the file:
	file, e := os.OpenFile("_go.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if e != nil {
		log.Fatal("Failed to open/create log file:", e)
	}
	defer file.Close()
	info.SetOutput(file)
	err.SetOutput(file)

	info.Println("Message for info logger")
	err.Println("Message for error logger")
}

var info = log.New(os.Stdout, "I ", log.Ldate|log.Ltime|log.Lmicroseconds| /*log.Lmsgprefix|*/ log.Lshortfile)
var err = log.New(os.Stdout, "E ", log.Ldate|log.Ltime|log.Lmicroseconds| /*log.Lmsgprefix|*/ log.Lshortfile)
