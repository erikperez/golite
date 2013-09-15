// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"path/filepath"
	"fmt"
	"io/ioutil"
	"net/http"
	"flag"
	"strconv"
)

var lenPath = len(rootDirectory)

func viewHandler(w http.ResponseWriter, r *http.Request) {

	filename := r.URL.Path[lenPath:]
	absoluteFileName := rootDirectory+"/"+filename
	fmt.Printf("Received request for path: "+absoluteFileName + "\n")
	body, err := ioutil.ReadFile(absoluteFileName)

	if (body == nil){
		fmt.Fprintf(w, "<html><head>Oops..the string has snapped!</head><body><div>%s</div></body></html>", err)
	}
	detectExtensionAndSetResponseHeader(w, absoluteFileName)
	fmt.Fprintf(w, "%s", body)
}

func detectExtensionAndSetResponseHeader(response http.ResponseWriter, absoluteFileName string) {
	extension := filepath.Ext(absoluteFileName)
	setContenType(response, contentTypes[extension])
}

var contentTypes = map[string]string{
    //Images
    ".png"	: 	"image/png",
    ".jpg"	: 	"image/jpeg",
    ".gif"	: 	"image/gif",
    ".tiff"	: 	"image/",
    ".jpeg"	: 	"image/jpeg",
    //Text
    ".css"	: 	"text/css",
    ".html"	: 	"text/html",
    ".htm" 	: 	"text/html",
    ".js"	: 	"application/javascript",
    ".json"	:	"application/json",
    ".exe"	: 	"application/octet-stream",
    ".zip"	:	"application/x-zip-compressed",
}


func setContenType(response http.ResponseWriter, contentType string) {
	response.Header().Set("Content-Type", contentType)
}



var port int
var rootDirectory string

func readFlags() {
	flag.IntVar(&port, "port", 8080, "The port the server is listening on")
	flag.StringVar(&rootDirectory, "root", "root", "The web root directory")
	flag.Parse()
}

func main() {
	readFlags()
	http.HandleFunc("/", viewHandler)
	fmt.Printf("Added root directory \n")
	fmt.Printf("Now listening for port "+strconv.Itoa(port) + "\n")
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
	
}			

