package trace

type SQL struct {
	Timestamp   string  `json:"timestamp"`     // Time, format: 2006-01-02 15:04:05
	Stack       string  `json:"stack"`         // file address and line number
	SQL         string  `json:"sql"`           // SQL statement
	Rows        int64   `json:"rows_affected"` // The number of rows affected
	CostSeconds float64 `json:"cost_seconds"`  // execution time (in seconds)
}
