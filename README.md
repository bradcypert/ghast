# ghast

[![Build Status](https://travis-ci.org/bradcypert/ghast.svg?branch=master)](https://travis-ci.org/bradcypert/ghast)

Ghast is an "All in one" toolkit for building reliable Go web applications.

Whether you're building an API, website, or something a little more exotic (still over HTTP/HTTPS protocols)
Ghast has got your back. Ghast is a collection of common functionality that I've extracted and refactored from several different Golang APIs and takes inspiration from classics such as "Rails" and "Laravel".

Here's what you should know about Ghast:

1. It's lightweight. The framework should be seen as bare helpers to the standard library.
2. We support any database that Gorm supports because Ghast uses Gorm.
3. Ghast currently follows the MVC paradigm closely. If this isn't your cup of tea, you MAY still benefit from Ghast, but it really works best when all pieces play together nicely.
4. Ghast ships with a CLI to help you generate your ghast controllers, models, and much more!

# Ghast CLI

Still here? Great! You can install Ghast's CLI by running:

```bash
go get github.com/bradcypert/ghast
```

### Create a new Ghast project

```bash
ghast new MyProject
```

### Create a new Controller

```bash
ghast make controller UsersController
```

# GhastRouter

### Setting up route bindings in Ghast

GhastRouter has support for all the common HTTP verbs and allows you to setup your routes based off of HTTP Verb and path.

```go
import (
	"fmt"
	"log"
	"net/http"

	ghastRouter "github.com/bradcypert/ghast/pkg/router"
)

func main() {
	router := ghastRouter.Router{}

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Hello from Ghast!")
	})

	s := router.DefaultServer()
	log.Fatal(s.ListenAndServe())
}
```

### Path Variables

GhastRouter also allows you to specify flexible route pattern that can be used to retrieve values from the route.

```go
router.Get("/:name", func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    fmt.Fprint(w, "Hello "+r.Context().Value("name").(string))
})
```

### Using Middleware with GhastRouter

GhastRouter supports two different styles of middleware.

Global Middleware is triggered on any route match before we delegate to your handler func.
Route specific middleware is triggered on the route it is specified on before we delegate to your handler func.

Global middleware is triggered before route specific middleware.

```go
router := ghastRouter.Router{}

middleware := []ghastRouter.MiddlewareFunc{
    func(rw *http.ResponseWriter, req *http.Request) {
        fmt.Println("Incoming Request: " + req.URL.String())
    },
}

router.AddMiddleware(middleware)
```

### Using Route Specific Middleware

GhastRouter also supports middleware for a single route.

```go
middleware := []ghastRouter.MiddlewareFunc{
    func(rw *http.ResponseWriter, req *http.Request) {
        fmt.Println("Incoming Request: " + req.URL.String())
    },
}

// Notice the `.GetM` method? We use `GetM` instead of `Get` to allow us to add route specific middleware
router.GetM("/:name", func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    fmt.Fprint(w, "Hello "+r.Context().Value("name").(string))
}, middleware)
```

# GhastControllers

soon™

# GhastModels (IE: common Gorm helpers)

soon™

# GhastFactories

soon™

More documentation coming soon!
