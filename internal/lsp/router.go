package lsp

import (
	"encoding/json"
	"log"
)

func Route(method string, contents []byte) (interface{}, error) {

	log.Printf("Received msg with method: %s", method)
	if method == "initialize" {
		log.Printf("Received msg with contents: %s", contents)
	}
	switch method {
	case "initialize":
		var request InitializeRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			return nil, err
		}

		log.Printf("Connected to: %s %s",
			request.Params.ClientInfo.Name,
			request.Params.ClientInfo.Version)

		msg := NewInitializeResponse(request.ID)
		return msg, nil
	case "textDocument/didOpen":
		var request DidOpenTextDocumentNotification
		if err := json.Unmarshal(contents, &request); err != nil {
			return nil, err
		}

		log.Printf("Opened: %s", request.Params.TextDocument.URI)
		return nil, nil
	default:
		return nil, nil
	}
}
