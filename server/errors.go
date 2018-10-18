package server

import "encoding/json"

// Various error messages
var (
	ErrorInvalidSignature, _ = json.Marshal(map[string]string{
		"error": "Invalid signature",
	})
	ErrorInvalidPayload, _ = json.Marshal(map[string]string{
		"error": "Invalid payload",
	})
	ErrorInternalServer, _ = json.Marshal(map[string]string{
		"error": "Internal Server Error",
	})
)
