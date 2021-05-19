package trace

import (
	"crypto/rand"
	"encoding/hex"
	"io"
	"sync"
)

const Header = "TRACE-ID"

var _ T = (*Trace)(nil)

type T interface {
	i()
	ID() string
	WithRequest(req *Request) *Trace
	WithResponse(resp *Response) *Trace
	AppendDialog(dialog *Dialog) *Trace
	AppendSQL(sql *SQL) *Trace
	AppendRedis(redis *Redis) *Trace
	AppendGRPC(grpc *Grpc) *Trace
}

// Trace recorded parameters
type Trace struct {
	mux                sync.Mutex
	Identifier         string    `json:"trace_id"`             // link ID
	Request            *Request  `json:"request"`              // Request information
	Response           *Response `json:"response"`             // return information
	ThirdPartyRequests []*Dialog `json:"third_party_requests"` // Information about calling third party interface
	Debugs             []*Debug  `json:"debugs"`               // debugging information
	SQLs               []*SQL    `json:"sqls"`                 // executed SQL information
	Redis              []*Redis  `json:"redis"`                // Redis information executed
	GRPCs              []*Grpc   `json:"grpc"`                 // executed gRPC information
	Success            bool      `json:"success"`              // request result true or false
	CostSeconds        float64   `json:"cost_seconds"`         // execution time (in seconds)
}

// Request request information
type Request struct {
	TTL        string      `json:"ttl"`         // request timeout
	Method     string      `json:"method"`      // request method
	DecodedURL string      `json:"decoded_url"` // request address
	Header     interface{} `json:"header"`      // Request Header information
	Body       interface{} `json:"body"`        // Request Body information
}

// Response 响应信息
type Response struct {
	Header          interface{} `json:"header"`                      // Header information
	Body            interface{} `json:"body"`                        // Body information
	BusinessCode    int         `json:"business_code,omitempty"`     // business code
	BusinessCodeMsg string      `json:"business_code_msg,omitempty"` // prompt message
	HttpCode        int         `json:"http_code"`                   // HTTP status code
	HttpCodeMsg     string      `json:"http_code_msg"`               // HTTP status code information
	CostSeconds     float64     `json:"cost_seconds"`                // execution time (in seconds)
}

func New(id string) *Trace {
	if id == "" {
		buf := make([]byte, 10)
		io.ReadFull(rand.Reader, buf)
		id = string(hex.EncodeToString(buf))
	}

	return &Trace{
		Identifier: id,
	}
}

func (t *Trace) i() {}

// ID Unique identifier
func (t *Trace) ID() string {
	return t.Identifier
}

// WithRequest set request
func (t *Trace) WithRequest(req *Request) *Trace {
	t.Request = req
	return t
}

// WithResponse set response
func (t *Trace) WithResponse(resp *Response) *Trace {
	t.Response = resp
	return t
}

// AppendDialog safely append internal call procedure dialog
func (t *Trace) AppendDialog(dialog *Dialog) *Trace {
	if dialog == nil {
		return t
	}

	t.mux.Lock()
	defer t.mux.Unlock()

	t.ThirdPartyRequests = append(t.ThirdPartyRequests, dialog)
	return t
}

// AppendDebug append debug
func (t *Trace) AppendDebug(debug *Debug) *Trace {
	if debug == nil {
		return t
	}

	t.mux.Lock()
	defer t.mux.Unlock()

	t.Debugs = append(t.Debugs, debug)
	return t
}

// AppendSQL append SQL
func (t *Trace) AppendSQL(sql *SQL) *Trace {
	if sql == nil {
		return t
	}

	t.mux.Lock()
	defer t.mux.Unlock()

	t.SQLs = append(t.SQLs, sql)
	return t
}

// AppendRedis append Redis
func (t *Trace) AppendRedis(redis *Redis) *Trace {
	if redis == nil {
		return t
	}

	t.mux.Lock()
	defer t.mux.Unlock()

	t.Redis = append(t.Redis, redis)
	return t
}

// AppendGRPC append gRPC call information
func (t *Trace) AppendGRPC(grpc *Grpc) *Trace {
	if grpc == nil {
		return t
	}

	t.mux.Lock()
	defer t.mux.Unlock()

	t.GRPCs = append(t.GRPCs, grpc)
	return t
}
