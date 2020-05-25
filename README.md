# ghast

[![Build Status](https://travis-ci.org/bradcypert/ghast.svg?branch=master)](https://travis-ci.org/bradcypert/ghast)

Ghast is an "All in one" toolkit for building reliable Go web applications.

Whether you're building an API, website, or something a little more exotic, Ghast has got your back. Ghast is a collection of common functionality that I've extracted, refactored, and built upon from several different Golang APIs and takes inspiration from classics such as "Rails" and "Laravel".

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

# GhastApp

The Ghast framework ships encapsulated in the `GhastApp` struct. However, you can use individual pieces of Ghast if you would prefer. I've tried to make all of these pieces provide some value individually, but the router is the only real candidate for use outside of the framework at the moment. That being said, `GhastApp` takes care of setting up your application, dependency injection container, and router. If you're building a full-fledged application, I strongly recommend using the `GhastApp` as it is the intended way to use Ghast.

TLDR: If you're not sure if you want the `GhastApp` or `GhastRouter`, use the `GhastApp` unless you are specifically only wanting routing.

Using `ghast new MyProjectName` from the command line will generate a new main.go file for you that's setup to use `GhastApp`.

```go
package main

import (
	"fmt"
	"net/http"

	ghastApp "github.com/bradcypert/ghast/pkg/app"
	ghastRouter "github.com/bradcypert/ghast/pkg/router"
)

func main() {
	router := ghastRouter.Router{}

	// Want to use controllers instead? Try running "ghast make controller MyController" from your terminal
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Hello from Ghast!")
	})

	app := ghastApp.NewApp()
	app.SetRouter(router)
	app.Start()
}

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

GhastRouter also allows you to specify a flexible route pattern that can be used to retrieve values from the route. Check out Ghast's controllers if you'd like to find an even cleaner way to retrieve path parameters.

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

Controllers also provide a lot of functionality for working with the HTTP request and response writer objects.

```go
package controllers

import (
	"net/http"

	ghastController "github.com/bradcypert/ghast/pkg/controllers"
)

type TestController struct {
	ghastController.GhastController
}

func (c *TestController) Index(w http.ResponseWriter, r *http.Request) {
	// get user ID from path, /user/:user
	userId := c.PathParam("user").(string)
}
```

I strongly recommend checking out the GoDoc for Ghast to get a better understanding of what is possible with Ghast's controllers (and all of Ghast, for that matter).

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
