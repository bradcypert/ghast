package router

import (
	"net/http"
	"time"

	pathToRegexp "github.com/soongo/path-to-regexp"
)

// Binding maps a url pattern to an HTTP handler
type Binding struct {
	match   string
	handler func(http.ResponseWriter, *http.Request)
}

// Router struct models our routes to match off of and their behavior
type Router struct {
	gets     []Binding
	posts    []Binding
	patches  []Binding
	puts     []Binding
	deletes  []Binding
	options  []Binding
	heads    []Binding
	connects []Binding
	traces   []Binding
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

		// for _, g := range match.Groups() {
		// 	fmt.Printf("%q ", g.String())
		// }
		// fmt.Printf("%d %q\n", match.Index, match)

		if match != nil {
			binding.handler(rw, req)
		}
	}
}

// Get registers a new GET route with the router
func (r *Router) Get(route string, f func(http.ResponseWriter,
	*http.Request)) *Router {
	r.gets = append(r.gets, Binding{match: route, handler: f})
	return r
}

// Post registers a new POST route with the router
func (r *Router) Post(route string, f func(http.ResponseWriter,
	*http.Request)) *Router {
	r.posts = append(r.posts, Binding{match: route, handler: f})
	return r
}

// Put registers a new PUT route with the router
func (r *Router) Put(route string, f func(http.ResponseWriter,
	*http.Request)) *Router {
	r.puts = append(r.puts, Binding{match: route, handler: f})
	return r
}

// Patch registers a new PATCH route with the router
func (r *Router) Patch(route string, f func(http.ResponseWriter,
	*http.Request)) *Router {
	r.patches = append(r.patches, Binding{match: route, handler: f})
	return r
}

// Delete registers a new DELETE route with the router
func (r *Router) Delete(route string, f func(http.ResponseWriter,
	*http.Request)) *Router {
	r.deletes = append(r.deletes, Binding{match: route, handler: f})
	return r
}

// Options registers a new OPTIONS route with the router
func (r *Router) Options(route string, f func(http.ResponseWriter,
	*http.Request)) *Router {
	r.options = append(r.options, Binding{match: route, handler: f})
	return r
}

// Head registers a new HEAD route with the router
func (r *Router) Head(route string, f func(http.ResponseWriter,
	*http.Request)) *Router {
	r.heads = append(r.heads, Binding{match: route, handler: f})
	return r
}

// Trace registers a new TRACE route with the router
func (r *Router) Trace(route string, f func(http.ResponseWriter,
	*http.Request)) *Router {
	r.traces = append(r.traces, Binding{match: route, handler: f})
	return r
}

// DefaultServer is an optional method to help get a preconfigured server
// with the router bound as the handler and some sensible defaults
func (r Router) DefaultServer() *http.Server {
	return &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
