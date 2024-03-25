package main

import (
	"bufio"
	"io"
	"log"
	"os"

	"github.com/jborkows/tsf-lsp/internal/logs"
	"github.com/jborkows/tsf-lsp/internal/lsp"
	"github.com/jborkows/tsf-lsp/internal/rpc"
)

func main() {
	logger, err := logs.Initialize(logs.FileLogger("tsf-lsp.log"))
	if err != nil {
		panic(err)
	}
	defer logger.Close()

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	writer := os.Stdout
	state := lsp.ProvideState(func(v any) {
		writeResponse(writer, v)
	})

	defer state.Close()
	for scanner.Scan() {
		msg := scanner.Bytes()
		method, contents, err := rpc.DecodeMessage(msg)
		if err != nil {
			log.Printf("Got an error: %s", err)
			continue
		}

		handleMessage(writer, method, contents, state)
	}
}

func handleMessage(writer io.Writer, method string, contents []byte, state *lsp.State) {
	response, err := lsp.Route(method, contents, state)
	if err != nil {
		log.Printf("Got an error: %s", err)
		return
	}
	if response != nil {
		log.Printf("Sending response for %s", method)
		writeResponse(writer, response)
	}

}

func writeResponse(writer io.Writer, msg any) {
	reply := rpc.EncodeMessage(msg)
	writer.Write([]byte(reply))

}
