package router

import "net/http"

// Resource an interface that includes functionality intended to be
// used with fetching a list of items, getting a single item, creating
// a single item, updating a single item, and deleting a single item
type Resource interface {
	GetName() string
	Index(req *http.Request) (Response, error)
	Get(req *http.Request) (Response, error)
	Create(req *http.Request) (Response, error)
	Update(req *http.Request) (Response, error)
	Delete(req *http.Request) (Response, error)
}
