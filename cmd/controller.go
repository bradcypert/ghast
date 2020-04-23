package cmd

import (
	"fmt"
	"os"
	"text/template"

	"github.com/spf13/cobra"
)

type controllerOptions struct {
	Package string
	Name    string
}

var pkg string
var name string

// controllerCmd represents the controller command
var controllerCmd = &cobra.Command{
	Use:   "controller",
	Args:  cobra.MinimumNArgs(1),
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		options := controllerOptions{
			pkg,
			name,
		}

		t := template.Must(template.New("controller").Parse(controllerTemplate))
		os.Mkdir("controllers", 0777)
		f, err := os.Create(fmt.Sprintf("./controllers/%s.go", name))
		if err != nil {
			panic("Unable to create controller")
		}
		t.Execute(f, options)
		f.Close()
	},
}

func init() {
	makeCmd.AddCommand(controllerCmd)
	controllerCmd.Flags().StringVarP(&pkg, "package", "p", "controllers", "Package name")
}

var controllerTemplate = `
package {{.Package}}

type {{.Name}} struct {}

func (c *{{.Name}}) Index() int {
  	return c.list.Len()
}

func (c *{{.Name}}) Get() int {
	return c.list.Len()
}

func (c *{{.Name}}) Create() int {
	return c.list.Len()
}

func (c *{{.Name}}) Edit() int {
	return c.list.Len()
}

func (c *{{.Name}}) Update() int {
	return c.list.Len()
}

func (c *{{.Name}}) Delete() int {
	return c.list.Len()
}
`
