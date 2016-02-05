package store

import (
	"bufio"
	"bytes"
	"log"
	"net/http"
)

type Store struct {
	Requests []Entry
}

type Entry struct {
	Request http.Request
}

func (store *Store) SaveRequest(request *http.Request) {
	buf := new(bytes.Buffer)
	request.Write(buf)
	newReq, err := http.ReadRequest(bufio.NewReader(bytes.NewReader(buf.Bytes())))
	if err != nil {
		log.Fatal("Should never happen in store", err)
	}

	store.Requests = append(store.Requests, Entry{Request: *newReq})
}

func (store *Store) GetLatest() *http.Request {
	if len(store.Requests) < 1 {
		return nil
	}

	return &(store.Requests[len(store.Requests)-1].Request)
}

func (store *Store) Get(i int) *http.Request {
	if len(store.Requests) < 1 || i > len(store.Requests)-1 {
		return nil
	}
	return &(store.Requests[i].Request)
}
