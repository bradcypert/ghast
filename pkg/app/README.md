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