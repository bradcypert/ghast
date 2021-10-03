# Ghast's DI Container

Ghast ships with a rudimentary DI container. Future plans include expanding upon this DI container and ultimately running all of Ghast through the container. For now, you can work with your own DI container like so:

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