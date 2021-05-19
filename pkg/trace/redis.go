package trace

type Redis struct {
	Timestamp   string  `json:"timestamp"`       // Time, format: 2006-01-02 15:04:05
	Handle      string  `json:"handle"`          // Operation, SET/GET etc.
	Key         string  `json:"key"`             // Key
	Value       string  `json:"value,omitempty"` // Value
	TTL         float64 `json:"ttl,omitempty"`   // Timeout period (unit minutes)
	CostSeconds float64 `json:"cost_seconds"`    // execution time (in seconds)
}
