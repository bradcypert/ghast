# GhastControllers

Ghast's Controllers are extremely easy to use and provide a few minor-abstractions to help you write simpler and cleaner code.

To use a GhastController, you can either run `ghast make controller MyControllerNameHere` or create your own controller and embed the `GhastController` struct into it.

Ghast's controllers are a little different than the standard http.handlers that the router supports (although technically a controller can use those too!). These controllers are intended to be used with a signature like so: `func (c *TestController) Index(r *http.Request) (router.Response, error)`, however, binding these to a router requires a special step.

Namely, controllers that are built this way need to have their receiver functions passed to the router.RouteFunc to add support for the serveHTTP method thats necessary for our router (essentially making these conform to http.Handler from the standard library).

A simple example of a controller can be seen here:

```go
package controllers

import (
	"net/http"

	"github.com/bradcypert/ghast/pkg/controllers"
	"github.com/bradcypert/ghast/pkg/router"
)

type TestController struct {
	controllers.GhastController
}

type Thing struct {
	Foo string
}

func (c *TestController) Index(r *http.Request) (router.Response, error) {
	myStruct := Thing{Foo: "bar"}
	return router.Response{
		Body: myStruct
	}, nil
}
```

Controllers also provide a lot of functionality for working with the HTTP request and response writer objects.

```go
package controllers

import (
	"net/http"

	"github.com/bradcypert/ghast/pkg/controllers"
)

type TestController struct {
	controllers.GhastController
}

func (c *TestController) Index(r *http.Request) (router.Response, error) {
	// get user ID from path, /user/:user
	userId := c.PathParam("user").(string)
	
	// ...
}
```

I strongly recommend checking out the GoDoc for Ghast to get a better understanding of what is possible with Ghast's controllers (and all of Ghast, for that matter).
