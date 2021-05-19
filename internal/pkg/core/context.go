package core

import (
	"bytes"
	stdctx "context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"sync"

	"github.com/xinliangnote/go-gin-api/pkg/errno"
	"github.com/xinliangnote/go-gin-api/pkg/trace"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.uber.org/zap"
)

type HandlerFunc func(c Context)

type Trace = trace.T

const (
	_Alias            = "_alias_"
	_TraceName        = "_trace_"
	_LoggerName       = "_logger_"
	_BodyName         = "_body_"
	_PayloadName      = "_payload_"
	_GraphPayloadName = "_graph_payload_"
	_UserID           = "_user_id_"
	_UserName         = "_user_name_"
	_AbortErrorName   = "_abort_error_"
)

var contextPool = &sync.Pool{
	New: func() interface{} {
		return new(context)
	},
}

func newContext(ctx *gin.Context) Context {
	context := contextPool.Get().(*context)
	context.ctx = ctx
	return context
}

func releaseContext(ctx Context) {
	c := ctx.(*context)
	c.ctx = nil
	contextPool.Put(c)
}

var _ Context = (*context)(nil)

type Context interface {
	init()

	// ShouldBindQuery deserialize querystring
	// tag: `form:"xxx"` (Note: Do not write query)
	ShouldBindQuery(obj interface{}) error

	// ShouldBindPostForm Deserialize postform (querystring will be ignored)
	// tag: `form:"xxx"`
	ShouldBindPostForm(obj interface{}) error

	// ShouldBindForm deserialize querystring and postform at the same time;
	// When querystring and postform have the same fields, postform is used first.
	// tag: `form:"xxx"`
	ShouldBindForm(obj interface{}) error

	// ShouldBindJSON deserialize postjson
	// tag: `json:"xxx"`
	ShouldBindJSON(obj interface{}) error

	// ShouldBindURI deserialize the path parameter (for example, the routing path is /user/:name)
	// tag: `uri:"xxx"`
	ShouldBindURI(obj interface{}) error

	// Redirect redirect
	Redirect(code int, location string)

	// Trace get trace object
	Trace() Trace
	setTrace(trace Trace)
	disableTrace()

	// Logger get logger object
	Logger() *zap.Logger
	setLogger(logger *zap.Logger)

	// Payload return correctly
	Payload(payload interface{})
	getPayload() interface{}

	// GraphPayload GraphQL return value is different from api return structure
	GraphPayload(payload interface{})
	getGraphPayload() interface{}

	// HTML back to the interface
	HTML(name string, obj interface{})

	// AbortWithError error return
	AbortWithError(err errno.Error)
	abortError() errno.Error

	// Header get header object
	Header() http.Header
	// GetHeader get header
	GetHeader(key string) string
	// SetHeader set header
	SetHeader(key, value string)

	// UserID get userId in JWT
	UserID() int64
	setUserID(userID int64)

	// UserName get userName in JWT
	UserName() string
	setUserName(userName string)

	// Alias set routing alias for metrics uri
	Alias() string
	setAlias(path string)

	// RequestInputParams get all parameters
	RequestInputParams() url.Values
	// RequestPostFormParams  get PostForm parameters
	RequestPostFormParams() url.Values
	// Request get Request object
	Request() *http.Request
	// RawData get Request.Body
	RawData() []byte
	// Method get Request.Method
	Method() string
	// Host get Request.Host
	Host() string
	// Path get the requested path Request.URL.Path (without querystring)
	Path() string
	// URI Request.URL.RequestURI() after getting unescape
	URI() string
	// RequestContext get the requested context (when the client is closed, it will be automatically canceled)
	RequestContext() StdContext

	// ResponseWriter obtain a ResponseWriter object
	ResponseWriter() gin.ResponseWriter
}

type context struct {
	ctx *gin.Context
}

type StdContext struct {
	stdctx.Context
	Trace
	*zap.Logger
}

func (c *context) init() {
	body, err := c.ctx.GetRawData()
	if err != nil {
		panic(err)
	}

	c.ctx.Set(_BodyName, body)                                   // cache the body is used for trace
	c.ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body)) // re-construct req body
}

// ShouldBindQuery deserialize querystring
// tag: `form:"xxx"` (Note: Do not write query)
func (c *context) ShouldBindQuery(obj interface{}) error {
	return c.ctx.ShouldBindWith(obj, binding.Query)
}

// ShouldBindPostForm deserialize postform (querystring will be ignored)
// tag: `form:"xxx"`
func (c *context) ShouldBindPostForm(obj interface{}) error {
	return c.ctx.ShouldBindWith(obj, binding.FormPost)
}

// ShouldBindForm deserialize querystring and postform at the same time;
// When querystring and postform have the same field, postform is used first.
// tag: `form:"xxx"`
func (c *context) ShouldBindForm(obj interface{}) error {
	return c.ctx.ShouldBindWith(obj, binding.Form)
}

// ShouldBindJSON deserialize postjson
// tag: `json:"xxx"`
func (c *context) ShouldBindJSON(obj interface{}) error {
	return c.ctx.ShouldBindWith(obj, binding.JSON)
}

// ShouldBindURI deserialize the path parameter (for example, the routing path is /user/:name)
// tag: `uri:"xxx"`
func (c *context) ShouldBindURI(obj interface{}) error {
	return c.ctx.ShouldBindUri(obj)
}

// Redirect redirect
func (c *context) Redirect(code int, location string) {
	c.ctx.Redirect(code, location)
}

func (c *context) Trace() Trace {
	t, ok := c.ctx.Get(_TraceName)
	if !ok || t == nil {
		return nil
	}

	return t.(Trace)
}

func (c *context) setTrace(trace Trace) {
	c.ctx.Set(_TraceName, trace)
}

func (c *context) disableTrace() {
	c.setTrace(nil)
}

func (c *context) Logger() *zap.Logger {
	logger, ok := c.ctx.Get(_LoggerName)
	if !ok {
		return nil
	}

	return logger.(*zap.Logger)
}

func (c *context) setLogger(logger *zap.Logger) {
	c.ctx.Set(_LoggerName, logger)
}

func (c *context) getPayload() interface{} {
	if payload, ok := c.ctx.Get(_PayloadName); ok != false {
		return payload
	}
	return nil
}

func (c *context) Payload(payload interface{}) {
	c.ctx.Set(_PayloadName, payload)
}

func (c *context) getGraphPayload() interface{} {
	if payload, ok := c.ctx.Get(_GraphPayloadName); ok != false {
		return payload
	}
	return nil
}

func (c *context) GraphPayload(payload interface{}) {
	c.ctx.Set(_GraphPayloadName, payload)
}

func (c *context) HTML(name string, obj interface{}) {
	c.ctx.HTML(200, name+".html", obj)
}

func (c *context) Header() http.Header {
	header := c.ctx.Request.Header

	clone := make(http.Header, len(header))
	for k, v := range header {
		value := make([]string, len(v))
		copy(value, v)

		clone[k] = value
	}
	return clone
}

func (c *context) GetHeader(key string) string {
	return c.ctx.GetHeader(key)
}

func (c *context) SetHeader(key, value string) {
	c.ctx.Header(key, value)
}

func (c *context) UserID() int64 {
	val, ok := c.ctx.Get(_UserID)
	if !ok {
		return 0
	}

	return val.(int64)
}

func (c *context) setUserID(userID int64) {
	c.ctx.Set(_UserID, userID)
}

func (c *context) UserName() string {
	val, ok := c.ctx.Get(_UserName)
	if !ok {
		return ""
	}

	return val.(string)
}

func (c *context) setUserName(userName string) {
	c.ctx.Set(_UserName, userName)
}

func (c *context) AbortWithError(err errno.Error) {
	if err != nil {
		httpCode := err.GetHttpCode()
		if httpCode == 0 {
			httpCode = http.StatusInternalServerError
		}

		c.ctx.AbortWithStatus(httpCode)
		c.ctx.Set(_AbortErrorName, err)
	}
}

func (c *context) abortError() errno.Error {
	err, _ := c.ctx.Get(_AbortErrorName)
	return err.(errno.Error)
}

func (c *context) Alias() string {
	path, ok := c.ctx.Get(_Alias)
	if !ok {
		return ""
	}

	return path.(string)
}

func (c *context) setAlias(path string) {
	if path = strings.TrimSpace(path); path != "" {
		c.ctx.Set(_Alias, path)
	}
}

// RequestInputParams get all parameters
func (c *context) RequestInputParams() url.Values {
	_ = c.ctx.Request.ParseForm()
	return c.ctx.Request.Form
}

// RequestPostFormParams get PostForm parameters
func (c *context) RequestPostFormParams() url.Values {
	_ = c.ctx.Request.ParseForm()
	return c.ctx.Request.PostForm
}

// Request get Request
func (c *context) Request() *http.Request {
	return c.ctx.Request
}

func (c *context) RawData() []byte {
	body, ok := c.ctx.Get(_BodyName)
	if !ok {
		return nil
	}

	return body.([]byte)
}

// Method requested method
func (c *context) Method() string {
	return c.ctx.Request.Method
}

// Host requested host
func (c *context) Host() string {
	return c.ctx.Request.Host
}

// Path the requested path (without querystring)
func (c *context) Path() string {
	return c.ctx.Request.URL.Path
}

// URI uri after unescape
func (c *context) URI() string {
	uri, _ := url.QueryUnescape(c.ctx.Request.URL.RequestURI())
	return uri
}

// RequestContext (Package Trace + Logger) Get the requested context (when the client is closed, it will be automatically cancelled)
func (c *context) RequestContext() StdContext {
	return StdContext{
		//c.ctx.Request.Context(),
		stdctx.Background(),
		c.Trace(),
		c.Logger(),
	}
}

// ResponseWriter get ResponseWriter
func (c *context) ResponseWriter() gin.ResponseWriter {
	return c.ctx.Writer
}
