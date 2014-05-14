package main

import (
	"net/http"

	"github.com/hjr265/go-zrsc/zrsc"
)

func main() {
	err := http.ListenAndServe(":8080", http.StripPrefix("", http.FileServer(zrsc.HttpDir(""))))
	catch(err)
}

func catch(err error) {
	if err != nil {
		panic(err)
	}
}
