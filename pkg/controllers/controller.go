package controllers

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/CloudyKit/jet"
	ghastApp "github.com/bradcypert/ghast/pkg/app"
	ghastContainer "github.com/bradcypert/ghast/pkg/container"
)

// GhastController should be embedded into consumer controllers
// and provides helper functions for working with the responseWriter, etc.
type GhastController struct{}

// Container returns the DI container associated with the given controller/request pairing.
func (c GhastController) Container() *ghastContainer.Container {
	return ghastApp.AppContext.Value("ghast/container").(*ghastContainer.Container)
}

// Config gets a config value from the controller's container.
// Config keys map to YAML in a flattened dot structure prefixed by an @.
// For example:
// a:
//   b: "c"
// "c" can be retrieved via @a.b
// We can't guarantee the type, so we return interface{}
func (c GhastController) Config(key string) interface{} {
	return c.Container().Make(key)
}

// PathParam Get a Path Parameter from a given request and key
func (c GhastController) PathParam(r *http.Request, key string) interface{} {
	return r.Context().Value(key)
}

// PathParam Get a Path Parameter from a given request and key
// returns a list of strings as it supports multiple values for a given path param
func (c GhastController) QueryParam(r *http.Request, key string) []string {
	return r.URL.Query()[key]
}

// Unmarshal unmarshalls a request with a body into the provided struct
// returns an error or nil value depending on if the unmarshall succeeded or not.
func (c GhastController) Unmarshal(r *http.Request, s interface{}) error {
	body := r.Body

	payload, err := ioutil.ReadAll(body)

	if err != nil {
		return err
	}

	return json.Unmarshal(payload, s)
}

// View executes a view from the app templates
// returns a response with the body set to the template
// Feel free to modify the response object further before returning
// in your controller
func (c GhastController) View(name string, vars jet.VarMap, contextualData interface{}) (Response, error) {
	tmpl, err := ghastApp.GetApp(c.Container()).GetViewSet().GetTemplate(name)
	if err != nil {
		return Response{}, err
	}

	var b bytes.Buffer
	writer := bufio.NewWriter(&b)

	err = tmpl.Execute(writer, vars, contextualData)

	return Response{
		Body: b.Bytes(),
	}, err
}
