package main

import (
	//"bytes"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/jonfk/http-loupe/serialization"
	"github.com/jonfk/http-loupe/store"

	"gopkg.in/readline.v1"
)

var (
	server Server = Server{
		StoreLock: new(sync.RWMutex),
	}
)

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("serving on :8080")
	go func() {
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			log.Fatal(err)
		}
	}()
	InputReadline()
}

func handler(w http.ResponseWriter, r *http.Request) {
	//serialization.WriteToFile("temp.json", *r)
	server.StoreLock.Lock()
	server.Store.SaveRequest(r)
	server.StoreLock.Unlock()
	fmt.Fprintf(w, "ok printed\n")
}

type Server struct {
	StoreLock *sync.RWMutex
	Store     store.Store
}

func InputReadline() {
	var completer = readline.NewPrefixCompleter(
		readline.PcItem("say",
			readline.PcItem("hello"),
			readline.PcItem("bye"),
		),
		readline.PcItem("help"),
		readline.PcItem("ping"),
		readline.PcItem("print"),
	)

	rl, err := readline.NewEx(&readline.Config{
		Prompt:       "> ",
		AutoComplete: completer,
	})
	if err != nil {
		panic(err)
	}
	defer rl.Close()

	for {
		line_, err := rl.Readline()
		if err != nil { // io.EOF
			break
		}
		line := strings.Split(line_, " ")
		if len(line) < 1 {
			continue
		}
		switch line[0] {
		case "ping":
			fmt.Println("ping")
		case "say":
			fmt.Println("say ", line)
		case "save":
			if len(line) < 2 {
				server.StoreLock.RLock()
				lastReq := server.Store.GetLatest()
				server.StoreLock.RUnlock()

				serialization.WriteToFile("temp.json", *lastReq)
			}
		}
	}
}
