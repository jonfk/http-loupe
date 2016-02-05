package main

import (
	//"bytes"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/jonfk/http-loupe/serialization"

	"gopkg.in/readline.v1"
)

var (
	server *Server = NewServer()
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
	server.StoreLock.Lock()
	server.Store.SaveRequest(r)
	server.StoreLock.Unlock()
	fmt.Fprintf(w, "ok\n")
}

func InputReadline() {
	var completer = readline.NewPrefixCompleter(
		readline.PcItem("help"),
		readline.PcItem("print"),
		readline.PcItem("list"),
		readline.PcItem("save"),
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
		line_ = strings.TrimRight(line_, " ")
		line := strings.Split(line_, " ")
		if len(line) < 1 {
			continue
		}
		switch line[0] {
		case "help":
			fmt.Println("help msg")
		case "save":
			save(server, line)
		case "print":
			print(server, line)
		case "list":
			list(server)
		}
	}
}

func list(server *Server) {
	allReqs := server.GetAllReqs()
	for i := range allReqs {
		fmt.Printf("%d : %v\n", i, allReqs[i])
	}
}

func print(server *Server, line []string) {
	var req *http.Request
	if len(line) < 2 {
		req = server.GetLatestReq()
	} else {
		i, err := strconv.Atoi(line[1])
		if err != nil {
			req = server.GetLatestReq()
		} else {
			req = server.GetReq(i)
		}
	}
	if req == nil {
		return
	}
	json, err := serialization.SerializeToJson(req)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(json))
}

func save(server *Server, line []string) {
	var req *http.Request
	if len(line) < 2 {
		req = server.GetLatestReq()
	} else {
		i, err := strconv.Atoi(line[1])
		if err != nil {
			req = server.GetLatestReq()
		} else {
			req = server.GetReq(i)
		}
	}
	if req == nil {
		return
	}
	err := serialization.WriteToFile("temp.json", req)
	if err != nil {
		fmt.Printf("[ERROR] %v\n", err)
	}
}
