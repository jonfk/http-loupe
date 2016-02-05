package main

import (
	"net/http"
	"sync"

	"github.com/jonfk/http-loupe/serialization"
	"github.com/jonfk/http-loupe/store"
)

type Server struct {
	// Inmemory store
	Store     *store.Store
	StoreLock *sync.RWMutex
	Options   *Options
}

type Options struct {
	PrintEveryRequest bool
}

func NewServer() *Server {
	return &Server{
		Store:     new(store.Store),
		StoreLock: new(sync.RWMutex),
		Options:   new(Options),
	}
}

func (s *Server) GetLatestReq() *http.Request {
	server.StoreLock.RLock()
	lastReq := server.Store.GetLatest()
	server.StoreLock.RUnlock()
	return lastReq
}

func (s *Server) GetReq(i int) *http.Request {
	server.StoreLock.RLock()
	req := server.Store.Get(i)
	server.StoreLock.RUnlock()
	return req
}

func (s *Server) GetAllReqs() []serialization.Request {
	server.StoreLock.RLock()
	copied := append([]store.Entry(nil), server.Store.Requests...)
	server.StoreLock.RUnlock()

	var result []serialization.Request
	for i := range copied {
		req, err := serialization.ToRequest(&copied[i].Request)
		if err != nil {
			continue
		}
		result = append(result, req)
	}

	return result
}
