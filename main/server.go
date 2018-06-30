package main

import (
	"net/http"
	"io"
	"awesomeProject/block"
	"encoding/json"
)

var chain *block.Chain

func run() {
	http.HandleFunc("/get", func(writer http.ResponseWriter, request *http.Request) {
		bytes, e := json.Marshal(chain)
		if e != nil {

		}

		io.WriteString(writer, string(bytes))
	})

	http.HandleFunc("/write", func(writer http.ResponseWriter, request *http.Request) {
		get := request.URL.Query().Get("data")
		chain.SendData(get)

		bytes, e := json.Marshal(chain)
		if e != nil {

		}

		io.WriteString(writer, string(bytes))
	})

	http.ListenAndServe("localhost:8888", nil)
}

func main() {
	chain = block.NewBlockChain()
	run()

}
