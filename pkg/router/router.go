package router

import (
    "context"
    "net/http"
    "time"

    ghastContainer "github.com/bradcypert/ghast/pkg/container"
    pathToRegexp "github.com/soongo/path-to-regexp"
)

// MiddlewareFunc is a functional alias to signify the signature of a middleware anonymous function.
type MiddlewareFunc = func(*http.ResponseWriter, *http.Request)

// Binding maps a url pattern to an HTTP handler
type Binding struct {
    match       string
    handler     func(http.ResponseWriter, *http.Request)
    middlewares []MiddlewareFunc
}

// Router struct models our routes to match off of and their behavior
type Router struct {
    container   *ghastContainer.Container
    gets        []Binding
    posts       []Binding
    patches     []Binding
    puts        []Binding
    deletes     []Binding
    options     []Binding
    heads       []Binding
    connects    []Binding
    traces      []Binding
    middlewares []MiddlewareFunc
}

// AddMiddleware registers a new global middleware within the router
// global middlewares will be called on any route that matches a route binding
func (r *Router) AddMiddleware(fn []MiddlewareFunc) *Router {
    r.middlewares = append(r.middlewares, fn...)
    return r
}

// SetDIContainer sets up the DI container that this router will use.
// The provided DI container will be available in all controllers via context (or controller helpers)
func (r *Router) SetDIContainer(c *ghastContainer.Container) {
    r.container = c
}

// ServeHTTP allows the router to adhere to the handler func requirements
// Delegates requests to provided routes
func (r Router) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
    var bindings []Binding

    switch req.Method {
    case "GET":
        bindings = r.gets
        break
    case "POST":
        bindings = r.posts
        break
    case "PUT":
        bindings = r.puts
        break
    case "PATCH":
        bindings = r.patches
        break
    case "DELETE":
        bindings = r.deletes
        break
    case "OPTIONS":
        bindings = r.options
        break
    case "HEAD":
        bindings = r.heads
        break
    case "TRACE":
        bindings = r.traces
        break
    case "CONNECT":
        bindings = r.connects
        break
    }

    for _, binding := range bindings {
        var tokens []pathToRegexp.Token
        regexp := pathToRegexp.Must(pathToRegexp.PathToRegexp(binding.match, &tokens, nil))

        match, _ := regexp.FindStringMatch(req.URL.String())

        if match != nil {

            // Run global middlewares
            for _, m := range r.middlewares {
                m(&rw, req)
            }

            // Run route specific middlewares
            for _, m := range binding.middlewares {
                m(&rw, req)
            }

            ctx := req.Context()
            for _, g := range match.Groups() {
                if len(tokens) >= match.Index && len(tokens) > 0 {
                    key := tokens[match.Index].Name
                    value := g.String()
                    ctx = context.WithValue(ctx, key, value)
                }
            }
            r := req.WithContext(ctx)
            binding.handler(rw, r)
            break
        }
    }
}

// PathParam Get a Path Parameter from a given request and key
func (r *Router) PathParam(req *http.Request, key string) interface{} {
    return req.Context().Value(key)
}

// Get registers a new GET route with the router
func (r *Router) Get(route string, f func(http.ResponseWriter,
    *http.Request)) *Router {
    r.gets = append(r.gets, Binding{match: route, handler: f})
    return r
}

// GetM registers a new GET route with the router and wires up the given middleware for that route only
func (r *Router) GetM(route string, f func(http.ResponseWriter,
    *http.Request), middleware []MiddlewareFunc) *Router {
    r.gets = append(r.gets, Binding{match: route, handler: f, middlewares: middleware})
    return r
}

// Post registers a new POST route with the router
func (r *Router) Post(route string, f func(http.ResponseWriter,
    *http.Request)) *Router {
    r.posts = append(r.posts, Binding{match: route, handler: f})
    return r
}

// PostM registers a new POST route with the router and wires up the given middleware for that route only
func (r *Router) PostM(route string, f func(http.ResponseWriter,
    *http.Request), middleware []MiddlewareFunc) *Router {
    r.posts = append(r.posts, Binding{match: route, handler: f, middlewares: middleware})
    return r
}

// Put registers a new PUT route with the router
func (r *Router) Put(route string, f func(http.ResponseWriter,
    *http.Request)) *Router {
    r.puts = append(r.puts, Binding{match: route, handler: f})
    return r
}

// PutM registers a new PUT route with the router and wires up the given middleware for that route only
func (r *Router) PutM(route string, f func(http.ResponseWriter,
    *http.Request), middleware []MiddlewareFunc) *Router {
    r.puts = append(r.puts, Binding{match: route, handler: f, middlewares: middleware})
    return r
}

// Patch registers a new PATCH route with the router
func (r *Router) Patch(route string, f func(http.ResponseWriter,
    *http.Request)) *Router {
    r.patches = append(r.patches, Binding{match: route, handler: f})
    return r
}

// PatchM registers a new PATCH route with the router and wires up the given middleware for that route only
func (r *Router) PatchM(route string, f func(http.ResponseWriter,
    *http.Request), middleware []MiddlewareFunc) *Router {
    r.patches = append(r.patches, Binding{match: route, handler: f, middlewares: middleware})
    return r
}

// Delete registers a new DELETE route with the router
func (r *Router) Delete(route string, f func(http.ResponseWriter,
    *http.Request)) *Router {
    r.deletes = append(r.deletes, Binding{match: route, handler: f})
    return r
}

// DeleteM registers a new DELETE route with the router and wires up the given middleware for that route only
func (r *Router) DeleteM(route string, f func(http.ResponseWriter,
    *http.Request), middleware []MiddlewareFunc) *Router {
    r.deletes = append(r.deletes, Binding{match: route, handler: f, middlewares: middleware})
    return r
}

// Options registers a new OPTIONS route with the router
func (r *Router) Options(route string, f func(http.ResponseWriter,
    *http.Request)) *Router {
    r.options = append(r.options, Binding{match: route, handler: f})
    return r
}

// OptionsM registers a new OPTIONS route with the router and wires up the given middleware for that route only
func (r *Router) OptionsM(route string, f func(http.ResponseWriter,
    *http.Request), middleware []MiddlewareFunc) *Router {
    r.options = append(r.options, Binding{match: route, handler: f, middlewares: middleware})
    return r
}

// Head registers a new HEAD route with the router
func (r *Router) Head(route string, f func(http.ResponseWriter,
    *http.Request)) *Router {
    r.heads = append(r.heads, Binding{match: route, handler: f})
    return r
}

// HeadM registers a new HEAD route with the router and wires up the given middleware for that route only
func (r *Router) HeadM(route string, f func(http.ResponseWriter,
    *http.Request), middleware []MiddlewareFunc) *Router {
    r.heads = append(r.heads, Binding{match: route, handler: f, middlewares: middleware})
    return r
}

// Trace registers a new TRACE route with the router
func (r *Router) Trace(route string, f func(http.ResponseWriter,
    *http.Request)) *Router {
    r.traces = append(r.traces, Binding{match: route, handler: f})
    return r
}

// TraceM registers a new TRACE route with the router and wires up the given middleware for that route only
func (r *Router) TraceM(route string, f func(http.ResponseWriter,
    *http.Request), middleware []MiddlewareFunc) *Router {
    r.traces = append(r.traces, Binding{match: route, handler: f, middlewares: middleware})
    return r
}

// DefaultServer is an optional method to help get a preconfigured server
// with the router bound as the handler and some sensible defaults
func (r Router) DefaultServer() *http.Server {
    return &http.Server{
        Addr:           ":9000",
        Handler:        r,
        ReadTimeout:    10 * time.Second,
        WriteTimeout:   10 * time.Second,
        MaxHeaderBytes: 1 << 20,
    }
}
