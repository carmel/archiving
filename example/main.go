//go:generate archiving -src=./public

package main

import (
	"log"
	"net/http"

	_ "github.com/carmel/archiving/example/archiving"
	"github.com/carmel/archiving/fs"
)

// Before buildling, run go generate.
// Then, run the main program and visit http://localhost:8080/public/hello.txt
func main() {
	archivingFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(archivingFS)))
	http.ListenAndServe(":8080", nil)
}
