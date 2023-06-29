package main

import (
	"net/http"

	"github.com/EugeneGpil/getFormattedUrls"
)

func main() {
	// create handler for route
	handler := func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello World!\n"))
	}

	mux := http.NewServeMux()

	// should be slashes here?
	//     V               V
	url := "get/hello/world"

	url1, url2 := getFormattedUrls.Run(url)

	//                                           V
	// for "http://localhost:8080/get/hello/world"
	mux.HandleFunc(url1, handler)

	//                                           V
	// for "http://localhost:8080/get/hello/world/"
	mux.HandleFunc(url2, handler)

	http.ListenAndServe(":8080", mux)
}
