// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const lenPath = len("/root/")

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[lenPath:]
	filename := title + ".txt"
	body, err := ioutil.ReadFile("root/"+filename)

	if (body == nil){
		fmt.Fprintf(w, "<div>%s</div>", err)
	}

	fmt.Fprintf(w, "<div>%s</div>", body)
}


type Configuration struct{
	Port int
	RootDirectory string
}

func main() {

	

	http.HandleFunc("/root/", viewHandler)
	http.ListenAndServe(":8080", nil)
}			

