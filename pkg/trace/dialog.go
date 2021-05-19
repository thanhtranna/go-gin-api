package trace

import "sync"

var _ D = (*Dialog)(nil)

type D interface {
	i()
	AppendResponse(resp *Response)
}

// Dialog internally calls the session information of the other party's interface; there will be a retry operation when it fails, so the response will be multiple times.
type Dialog struct {
	mux         sync.Mutex
	Request     *Request    `json:"request"`      // Request information
	Responses   []*Response `json:"responses"`    // return information
	Success     bool        `json:"success"`      // Whether it succeeded, true or false
	CostSeconds float64     `json:"cost_seconds"` // execution time (in seconds)
}

func (d *Dialog) i() {}

// AppendResponse additional response information by transfer
func (d *Dialog) AppendResponse(resp *Response) {
	if resp == nil {
		return
	}

	d.mux.Lock()
	d.Responses = append(d.Responses, resp)
	d.mux.Unlock()
}
