package router

import (
	"net/http"
)

// TODO: As of yet in this refactor do I know why these exists in this way and
// are named as such
type ChainHandler struct {
	Middlewares Middlewares
	Endpoint    http.Handler
	chain       http.Handler
}

// TODO: Why does Chain(middlewares...) return middlewares? this is chaining
// them?
func Chain(middlewares ...func(http.Handler) http.Handler) Middlewares {
	return Middlewares(middlewares)
}

// TODO: So chain is just the chain of itself and endpoint, then why is this
// needed? can we not just have this be the output of a function? And what can
// we do to make the router change-able over the course of the server being up
// for dev and non-dev reasons

// TODO: If I modify the handler here to take our new route type that holds more
// information and passes through or works with the
func (self Middlewares) Handler(h http.Handler) http.Handler {
	return &ChainHandler{
		Middlewares: self,
		Endpoint:    h,
		chain:       chain(self, h),
	}
}

// NOTE: This is different from above because it takes http.HandlerFunc instead
// of http's http.Handler // THOUGH it could be combined with a clever itnerface
// or taking in interface{} -- thinhk of any other ways?
func (self Middlewares) HandlerFunc(h http.HandlerFunc) http.Handler {
	return &ChainHandler{
		Middlewares: self,
		Endpoint:    h,
		chain:       chain(self, h),
	}
}

// When I figure out what this is doing I will really get this submodule
func (self *ChainHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	self.chain.ServeHTTP(w, r)
}

// so above bottom chain is chain(self (the self middlewares), h the endpoint)
// which is different than the field chain; quite confusing stufff seems
// uncessarily so to obstuficate or some shit
func chain(middlewares []func(http.Handler) http.Handler, endpoint http.Handler) http.Handler {
	if len(middlewares) == 0 {
		return endpoint
	}

	// Wrap the end handler with the middleware chain
	// so grabbing last, then (endpoint) on it, bc its a slice of functions
	//   SO we are running the function (last first), then we iterate through
	//   everything but the last backwards !! wtf lol
	// *first idea*  why not put it in different order so we dont have to do this?
	// *second* make this not stupid hard to read for no fucking reason
	// *third*
	h := middlewares[len(middlewares)-1](endpoint)
	for i := len(middlewares) - 2; i >= 0; i-- {
		h = middlewares[i](h)
	}

	return h
}
