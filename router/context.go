package router

import (
	"context"
	"net"
	"net/http"
	"strconv"
	"strings"
)

// TODO: I absolutely! want .Get(":slug").Int() or .Bool() etc
// slug := r.URL.Query().Get(":slug")

var (
	RouteCtxKey = &contextKey{"RouteContext"}
)

type Context struct {
	// Public
	Routes        Routes
	RoutePath     string
	RouteMethod   string
	RoutePatterns []string
	URLParams     RouteParams
	// Private
	routePattern     string
	routeParams      RouteParams
	methodNotAllowed bool
}

func NewRouteContext() *Context { return &Context{} }

func (self *Context) Reset() {
	self.Routes = nil
	self.RoutePath = ""
	self.RouteMethod = ""
	self.routePattern = ""
	self.RoutePatterns = self.RoutePatterns[:0]
	self.URLParams.Keys = self.URLParams.Keys[:0]
	self.URLParams.Values = self.URLParams.Values[:0]

	self.routeParams.Keys = self.routeParams.Keys[:0]
	self.routeParams.Values = self.routeParams.Values[:0]
	self.methodNotAllowed = false
}

// TODO: No create a URLParam type and then build out a Param.Int() etc. This
// will give us more control and clean up actions by providing the adequate
// endpoint ALSO it keeps the logic inline with choices made in cli
type Param struct {
	Name  string
	value string
}

func (self Param) Value() string  { return self.value }
func (self Param) String() string { return self.value }

// move data to its own subpackage and use it here again
func (self Param) Bool() bool {
	// data.IsTrue(self.value)
	for _, trueString := range []string{"t", "true", "1"} { // "y", "yes", "on"}
		if strings.ToLower(self.value) == trueString {
			return true
		}
	}
	return false
}

func (self Param) Int() int {
	intValue, err := strconv.Atoi(self.value)
	if err != nil {
		return 0
	} else {
		return intValue
	}
}

type Params []*Param

// TODO: Now we can create methods for each singular and plural types to
// encapsulate logic in the way that is most similar to our choices in cli which
// worked well but is also keeping the libraries consistent
// TODO: And why not just Param?
func (self *Context) URLParam(key string) string {
	for k := len(self.URLParams.Keys) - 1; k >= 0; k-- {
		if self.URLParams.Keys[k] == key {
			return self.URLParams.Values[k]
		}
	}
	return ""
}

func (self *Context) RoutePattern() string {
	routePattern := strings.Join(self.RoutePatterns, "")
	return strings.Replace(routePattern, "/*/", "/", -1)
}

func RouteContext(ctx context.Context) *Context {
	return ctx.Value(RouteCtxKey).(*Context)
}

func URLParam(r *http.Request, key string) string {
	if rctx := RouteContext(r.Context()); rctx != nil {
		return rctx.URLParam(key)
	}
	return ""
}

func URLParamFromCtx(ctx context.Context, key string) string {
	if rctx := RouteContext(ctx); rctx != nil {
		return rctx.URLParam(key)
	}
	return ""
}

// TODO: HATE THIS, use other data type
type RouteParams struct {
	Keys   []string
	Values []string
}

// TODO: We have a key/value store here, this may be a place of optimization
// TODO: This is quite clever, but is it too clever, and if not learn it lol
// TODO: NOT A FAN OF THIS! Pretty sure you can just switch it to non pointer
func (self *RouteParams) Add(key, value string) {
	(*self).Keys = append((*self).Keys, key)
	(*self).Values = append((*self).Values, value)
}

// TODO: When we understand this we get it
func ServerBaseContext(baseCtx context.Context, h http.Handler) http.Handler {
	fn := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		baseCtx := baseCtx
		if v, ok := ctx.Value(http.ServerContextKey).(*http.Server); ok {
			baseCtx = context.WithValue(baseCtx, http.ServerContextKey, v)
		}
		if v, ok := ctx.Value(http.LocalAddrContextKey).(net.Addr); ok {
			baseCtx = context.WithValue(baseCtx, http.LocalAddrContextKey, v)
		}

		h.ServeHTTP(w, r.WithContext(baseCtx))
	})
	return fn
}

type contextKey struct {
	name string
}

func (self *contextKey) String() string {
	return "router context value " + self.name
}
