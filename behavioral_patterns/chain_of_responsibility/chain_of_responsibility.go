package chain_of_responsibility

import "encoding/json"

type HttpCtx struct {
	handles []func(*HttpCtx)
	next    int
	status  int
	resp    string
}

func newHttpCtx() *HttpCtx {
	return &HttpCtx{
		handles: make([]func(*HttpCtx), 0),
		next:    0,
	}
}

func (c *HttpCtx) Next() {
	if c.next == len(c.handles) {
		return
	}
	c.next++
	c.handles[c.next-1](c)
}

func (c *HttpCtx) Resp(status int, resp interface{}) {
	marshal, _ := json.Marshal(resp)
	c.resp = string(marshal)
	c.status = status
}

type Router struct {
	ctx  *HttpCtx
	path string
}

func NewRouter() *Router {
	return &Router{
		ctx:  newHttpCtx(),
		path: "",
	}
}

func (r *Router) Use(handle func(*HttpCtx)) {
	r.ctx.handles = append(r.ctx.handles, handle)
}

func (r *Router) Get(path string, handle func(ctx *HttpCtx)) {
	r.path = path
	r.ctx.handles = append(r.ctx.handles, handle)
}

func (r *Router) Run() {
	r.ctx.Next()
	println(r.ctx.status, r.ctx.resp)
}
