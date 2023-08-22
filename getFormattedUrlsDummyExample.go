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

	urls := getFormattedUrls.Run(url)

	//                                           V
	// for "http://localhost:8080/get/hello/world"
	mux.HandleFunc(urls[0], handler)

	//                                           V
	// for "http://localhost:8080/get/hello/world/"
	mux.HandleFunc(urls[1], handler)

	http.ListenAndServe(":8080", mux)
}
