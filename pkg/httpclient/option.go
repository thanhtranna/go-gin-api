package httpclient

import (
	"sync"
	"time"

	"github.com/xinliangnote/go-gin-api/pkg/trace"

	"go.uber.org/zap"
)

var (
	cache = &sync.Pool{
		New: func() interface{} {
			return &option{
				header: make(map[string][]string),
			}
		},
	}
)

// Mock define interface mock data
type Mock func() (body []byte)

// Option custom setting http request
type Option func(*option)

type option struct {
	ttl         time.Duration
	header      map[string][]string
	trace       *trace.Trace
	dialog      *trace.Dialog
	logger      *zap.Logger
	retryTimes  int
	retryDelay  time.Duration
	retryVerify RetryVerify
	alarmTitle  string
	alarmObject AlarmObject
	alarmVerify AlarmVerify
	mock        Mock
}

func (o *option) reset() {
	o.ttl = 0
	o.header = make(map[string][]string)
	o.trace = nil
	o.dialog = nil
	o.logger = nil
	o.retryTimes = 0
	o.retryDelay = 0
	o.retryVerify = nil
	o.alarmTitle = ""
	o.alarmObject = nil
	o.alarmVerify = nil
	o.mock = nil
}

func getOption() *option {
	return cache.Get().(*option)
}

func releaseOption(opt *option) {
	opt.reset()
	cache.Put(opt)
}

// WithTTL the longest execution time of this http request
func WithTTL(ttl time.Duration) Option {
	return func(opt *option) {
		opt.ttl = ttl
	}
}

// WithHeader set the http header, you can call multiple times to set multiple pairs of key-value
func WithHeader(key, value string) Option {
	return func(opt *option) {
		opt.header[key] = []string{value}
	}
}

// WithTrace set trace information
func WithTrace(t trace.T) Option {
	return func(opt *option) {
		if t != nil {
			opt.trace = t.(*trace.Trace)
			opt.dialog = new(trace.Dialog)
		}
	}
}

// WithLogger set up logger to print key logs
func WithLogger(logger *zap.Logger) Option {
	return func(opt *option) {
		opt.logger = logger
	}
}

// WithMock set mock data
func WithMock(m Mock) Option {
	return func(opt *option) {
		opt.mock = m
	}
}

// WithOnFailedAlarm set up alarm notification
func WithOnFailedAlarm(alarmTitle string, alarmObject AlarmObject, alarmVerify AlarmVerify) Option {
	return func(opt *option) {
		opt.alarmTitle = alarmTitle
		opt.alarmObject = alarmObject
		opt.alarmVerify = alarmVerify
	}
}

// WithOnFailedRetry failed to set up and try again
func WithOnFailedRetry(retryTimes int, retryDelay time.Duration, retryVerify RetryVerify) Option {
	return func(opt *option) {
		opt.retryTimes = retryTimes
		opt.retryDelay = retryDelay
		opt.retryVerify = retryVerify
	}
}
