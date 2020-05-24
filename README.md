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

Ghast's Controllers are extremely easy to use and provide a few minor-abstractions to help you write simpler and cleaner code.

To use a GhastController, you can either run `ghast make controller MyControllerNameHere` or create your own controller and embed the `GhastController` struct into it.

Ghast's controllers give you access to a few convenient reciever functions on the controller instance. For each of the common HTTP status codes, there are helper functions that can be leveraged. For example, using `Success` will set the status code to success and write out whatever you pass as the 2nd argument. If the 2nd argument is a struct, Ghast will go ahead and marshall that, set the relevant content-type header to application/json, and then write that to the responseWriter.

A simple example of a controller can be seen here:

```go
package controllers

import (
	"net/http"

	ghastController "github.com/bradcypert/ghast/pkg/controllers"
)

type TestController struct {
	ghastController.GhastController
}

type Thing struct {
	Foo string
}

func (c *TestController) Index(w http.ResponseWriter, r *http.Request) {
	myStruct := Thing{Foo: "bar"}
	c.Success(w, myStruct)
}

```

# Ghast's DI Container

Ghast ships with a simple, yet rudimentary DI container. Future plans include expanding upon this DI container and ultimately running all of Ghast through the container. For now, you can work with your own DI container like so:

```go
package container

import (
	"testing"
	ghastContainer "github.com/bradcypert/ghast/pkg/container"
)

type Foo interface {
	hasFoo() bool
}

type Bar struct {
	secretKey string
}

func (b Bar) hasFoo() bool {
	return false
}

func TestResponses(t *testing.T) {
	t.Run("Should bind to the container correctly", func(t *testing.T) {
		container := ghastContainer.NewContainer()

		container.Bind("SECRET_KEY", func(container Container) interface{} {
			return "ABC123"
		})
		container.Bind("Bar", func(container Container) interface{} {
			return Bar{
				container.Make("SECRET_KEY").(string),
			}
		})

		bar := container.Make("Bar").(Bar)
		if bar.secretKey != "ABC123" {
			t.Errorf("Bound bar does not have the correct secret key")
		}

		if bar.hasFoo() != false {
			t.Errorf("Bound bar does not have the correct hasFoo implementation")
		}
	})
}
```

This example looks like a test, and that's because it is! You'll be able to find good examples of how to implement all of ghast by looking at the test files that live alongside the source files!

You'll notice that, at least as of now, you will have to perform a type conversion, as the container currently returns an interface{}. It is suggested that the DI keys that you bind against help provide context to the underlying type, so that other developers aren't confused about what is coming out of your container when they call `Make`.

# GhastModels (IE: common Gorm helpers)

soon™

# GhastFactories

soon™

More documentation coming soon!
