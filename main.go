package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {

	r, err := http.Get("http://reddit.com/r/golang.json")
	if err != nil {
		log.Fatal(err)
	}
	if r.StatusCode != http.StatusOK {
		log.Fatal(r.Status)
	}

	_, err = io.Copy(os.Stdout, r.Body)
	if err != nil {
		log.Fatal(err)
	}

}
