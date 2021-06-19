package router

import (
	"context"
	"math"
	"net/http"
	"strings"
)

type Interceptor func(routerContext *SeqRouterContext)

type InterceptorGroup struct {
	path         string
	interceptors []Interceptor
}

type SeqRouter struct {
	groups []*InterceptorGroup
	coreFunc func(routerContext *SeqRouterContext) http.Handler
}

type SeqRouterContext struct {
	Rw http.ResponseWriter
	Req *http.Request
	ctx context.Context
	*InterceptorGroup
	current int8
}

const endIndex int8 = math.MaxInt8

func NewSeqRouter() *SeqRouter {
	return &SeqRouter{}
}

func (r *SeqRouter) Group(path string, interceptors ...Interceptor) *SeqRouter {
	exists := false

	for _, group := range r.groups {
		if path == group.path {
			exists = true
			break
		}
	}

	if exists {
		return r
	}

	group := &InterceptorGroup{
		path:         path,
		interceptors: interceptors,
	}
	r.groups = append(r.groups, group)
	return r
}

func (r *SeqRouter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newSeqRouterContext(w, req, r)

	if r.coreFunc != nil {
		c.interceptors = append(c.interceptors, func(c *SeqRouterContext) {
			r.coreFunc(c).ServeHTTP(w, req)
		})
	}

	c.Reset()
	c.Next()
}

func newSeqRouterContext(w http.ResponseWriter, r *http.Request, router *SeqRouter) *SeqRouterContext {
	candGroup := &InterceptorGroup{}
	matchLen := 0

	for _, group := range router.groups {
		if strings.HasPrefix(r.RequestURI, group.path) {
			pathLen := len(group.path)

			if pathLen > matchLen {
				matchLen = pathLen
				candGroup = group
			}
		}
	}

	c := &SeqRouterContext{
		Rw:               w,
		Req:              r,
		InterceptorGroup: candGroup,
		ctx:              r.Context(),
	}
	c.Reset()
	return c
}

func (c *SeqRouterContext) Get(key interface{}) interface{} {
	return c.ctx.Value(key)
}

func (c *SeqRouterContext) Set(key, value interface{}) {
	c.ctx = context.WithValue(c.ctx, key, value)
}

func (c *SeqRouterContext) Next() {
	c.current++
	for c.current < int8(len(c.interceptors)) {
		c.interceptors[c.current](c)
		c.current++
	}
}

func (c *SeqRouterContext) Finish() {
	c.current = endIndex
}

func (c *SeqRouterContext) IsFinish() bool {
	return c.current >= endIndex
}

func (c *SeqRouterContext) Reset() {
	c.current = -1
}
