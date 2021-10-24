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

GhastRouter also allows you to specify a flexible route pattern that can be used to retrieve values from the route.

```go
router.Get("/:name", func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    fmt.Fprint(w, "Hello "+router.PathParam(r, "name").(string))
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

### Resources

The GhastRouter module also exposes an interface for Resources (named `Resource`). This interface defines the expectations for a struct (usually a controller) that successfully implements a common series of request handlers. In the case of `Resource`, the struct needs to implement `Index` (intended to list all items requested), `Get` (intended to list a single item requested), `Update` (update a single item), `Delete` (delete a single item), and `Create` (create a single item).

When a struct successfully implements `Resource`, you can leverage the router's `.Resource` method to generate the routes for that struct.

```go
type UserController struct {
	/// Implement resource here
}

router.Resource("/v1/", UserController{})
```