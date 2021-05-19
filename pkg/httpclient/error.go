package httpclient

var _ ReplyErr = (*replyErr)(nil)

// ReplyErr error response, used to wrap the returned httpcode and body when resp.StatusCode != http.StatusOK.
type ReplyErr interface {
	error
	StatusCode() int
	Body() []byte
}

type replyErr struct {
	err        error
	statusCode int
	body       []byte
}

func (r *replyErr) Error() string {
	return r.err.Error()
}

func (r *replyErr) StatusCode() int {
	return r.statusCode
}

func (r *replyErr) Body() []byte {
	return r.body
}

func newReplyErr(statusCode int, body []byte, err error) ReplyErr {
	return &replyErr{
		statusCode: statusCode,
		body:       body,
		err:        err,
	}
}

// ToReplyErr try to convert err to ReplyErr
func ToReplyErr(err error) (ReplyErr, bool) {
	if err == nil {
		return nil, false
	}

	e, ok := err.(ReplyErr)
	return e, ok
}
