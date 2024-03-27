package lspserver

import (
	"encoding/json"
	"fmt"
	"log"

	. "github.com/jborkows/tsf-lsp/internal/lsp"
)

func Route(method string, contents []byte, state *State) (interface{}, error) {

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

		msg := NewInitializeResponse(response(request.Request))
		return msg, nil
	case "textDocument/didOpen":
		var request DidOpenTextDocumentNotification
		if err := json.Unmarshal(contents, &request); err != nil {
			return nil, err
		}

		state.OpenDocument(request.Params.TextDocument.URI, request.Params.TextDocument.Text)
		return nil, nil
	case "textDocument/didChange":
		var request TextDocumentDidChangeNotification
		if err := json.Unmarshal(contents, &request); err != nil {
			return nil, err
		}
		state.UpdateDocument(request.Params.TextDocument.URI, request.Params.ContentChanges[0].Text)
		return nil, nil
	case "textDocument/hover":
		var request HoverRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			return nil, err
		}
		msg := state.Hover(request.Params.TextDocument.URI, request.Params.Position)
		response := HoverResponse{
			Response: response(request.Request),
			Result:   msg,
		}
		return response, nil
	case "textDocument/definition":
		var request DefinitionRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			return nil, err
		}
		location, error := state.DefinitionLocation(request.Params.TextDocument.URI, request.Params.Position)
		if error != nil {
			return nil, fmt.Errorf("Error getting definition: %w", error)
		}

		msg := DefinitionResponse{
			Response: response(request.Request),
			Result:   location,
		}
		return msg, nil
	case "textDocument/completion":
		var request CompletionRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			return nil, err
		}
		completions, error := state.Completion(request.Params.TextDocument.URI, request.Params.Position)
		if error != nil {
			return nil, fmt.Errorf("Error getting completions: %w", error)
		}
		msg := CompletionResponse{
			Response: response(request.Request),
			Result:   completions,
		}
		return msg, nil
	case "textDocument/codeAction":
		var request CodeActionRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			return nil, err
		}
		actions, error := state.CodeActions(request.Params.Range.Start, request.Params.TextDocument.URI)
		if error != nil {
			return nil, fmt.Errorf("Error getting code actions: %w", error)
		}
		msg := TextDocumentCodeActionResponse{
			Response: response(request.Request),
			Result:   actions,
		}

		return msg, nil
	case "workspace/executeCommand":
		var request ExecuteCommandRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			return nil, err
		}
		log.Printf("Received command: %v ", request)
		log.Printf("Received command: %s -> %s", request.Params.Command, request.Params.Arguments)
		return nil, nil

	case "textDocument/documentColor":
		var request ColorRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			return nil, err
		}
		colors := state.Color(request.Params.TextDocument.URI)
		msg := ColorResponse{
			Response: response(request.Request),
			Result:   colors,
		}

		return msg, nil
	default:
		return nil, nil
	}
}

func response(request Request) Response {
	return Response{
		RPC: "2.0",
		ID:  &request.ID,
	}
}
