package api

// PutRequest represents the structure for PUT operation requests
type PutRequest struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Response represents the structure for API responses
type Response struct {
	Status  string `json:"status"`
	Key     string `json:"key,omitempty"`
	Value   string `json:"value,omitempty"`
	Message string `json:"message,omitempty"`
}
