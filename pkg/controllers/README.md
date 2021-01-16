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