package router

import "net/http"

// Resource an interface that includes functionality intended to be
// used with fetching a list of items, getting a single item, creating
// a single item, updating a single item, and deleting a single item
type Resource interface {
	GetName() string
	Index(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}
