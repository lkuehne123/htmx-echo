package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Hello, World!")
	srv := &http.Server{
		Addr:     "4000",
		ErrorLog: log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
		Handler:  Routes(),
	}
	srv.ListenAndServe()
}
