package main

import (
	"bytes"
	"fmt"
	//"github.com/davecgh/go-spew/spew"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("serving on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	//spew.Dump(r)

	fmt.Println("Body:")

	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	strBody := buf.String()
	fmt.Println(strBody)

	ioutil.WriteFile("temp.txt", buf.Bytes(), 0777)

	fmt.Fprintf(w, "ok printed")
}
