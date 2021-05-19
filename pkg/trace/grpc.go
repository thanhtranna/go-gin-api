package trace

import (
	"google.golang.org/grpc/metadata"
)

type Grpc struct {
	Timestamp   string                 `json:"timestamp"`             // Time, format: 2006-01-02 15:04:05
	Addr        string                 `json:"addr"`                  // address
	Method      string                 `json:"method"`                // Operation method
	Meta        metadata.MD            `json:"meta"`                  // Mate
	Request     map[string]interface{} `json:"request"`               // Request information
	Response    map[string]interface{} `json:"response"`              // Return information
	CostSeconds float64                `json:"cost_seconds"`          // execution time (in seconds)
	Code        string                 `json:"err_code,omitempty"`    // error code
	Message     string                 `json:"err_message,omitempty"` // error message
}
