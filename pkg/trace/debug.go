package trace

type Debug struct {
	Key         string      `json:"key"`          // mark
	Value       interface{} `json:"value"`        // value
	CostSeconds float64     `json:"cost_seconds"` // execution time (in seconds)
}
