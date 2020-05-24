package container

import (
	"testing"
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
		container := NewContainer()

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
