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