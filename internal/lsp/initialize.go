package lsp

type InitializeRequest struct {
	Request
	Params InitializeRequestParams `json:"params"`
}

type InitializeRequestParams struct {
	ClientInfo *ClientInfo `json:"clientInfo"`
	// ... there's tons more that goes here
}

type ClientInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type InitializeResponse struct {
	Response
	Result InitializeResult `json:"result"`
}

type InitializeResult struct {
	Capabilities ServerCapabilities `json:"capabilities"`
	ServerInfo   ServerInfo         `json:"serverInfo"`
}
type ExecuteCommandClientCapabilities struct {
	Commands []string `json:"commands"`
}

type ServerCapabilities struct {
	TextDocumentSync   int                              `json:"textDocumentSync"`
	HoverProvider      bool                             `json:"hoverProvider"`
	DefinitionProvider bool                             `json:"definitionProvider"`
	CodeActionProvider bool                             `json:"codeActionProvider"`
	CompletionProvider map[string]any                   `json:"completionProvider"`
	ExecuteCommand     ExecuteCommandClientCapabilities `json:"executeCommand"`
}

type ServerInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func NewInitializeResponse(id int) InitializeResponse {
	return InitializeResponse{
		Response: Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: InitializeResult{
			Capabilities: ServerCapabilities{
				TextDocumentSync:   1,
				HoverProvider:      true,
				DefinitionProvider: true,
				CodeActionProvider: true,
				CompletionProvider: map[string]any{},
				ExecuteCommand: ExecuteCommandClientCapabilities{
					Commands: []string{"some_command"},
				},
			},
			ServerInfo: ServerInfo{
				Name:    "tsf-lsp",
				Version: "0.0.1",
			},
		},
	}
}
