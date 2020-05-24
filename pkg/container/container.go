package container

// Container provides a DI container to help manage application dependencies
type Container struct {
	bindings map[string]interface{}
}

// NewContainer creates a new DI container
func NewContainer() Container {
	return Container{
		make(map[string]interface{}),
	}
}

// Bind binds the result of the provided function as the value to be found when making the given key.
func (c Container) Bind(key string, fn func(c Container) interface{}) {
	c.bindings[key] = fn(c)
}

// Make creates an instance of the provided interface
func (c Container) Make(key string) interface{} {
	return c.bindings[key]
}
