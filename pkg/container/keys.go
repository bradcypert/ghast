package container

// These keys are defined for use by the Ghast framework.
// Please do not use these keys when defining your own bindings
// in the container.
// Keys that begin with an @ are generated when parsing
// the YAML config for a Ghast application.
const (
	AppKey    = "ghast/app"
	RouterKey = "ghast/router"
	PortKey   = "@ghast.config.port"
)
