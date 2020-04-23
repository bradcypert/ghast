package cmd

import (
	"fmt"
	"os"
	"text/template"

	"github.com/spf13/cobra"
)

type factoryOptions struct {
	Package string
	Name    string
}

// factoryCmd represents the factory command
var factoryCmd = &cobra.Command{
	Use:   "factory",
	Args:  cobra.MinimumNArgs(1),
	Short: "Create a new factory",
	Long: `Create a new factory.
	
	Factories are used for generating data for your application.
	They are particularly useful during local development and testing,
	although there are also reasons that you may want to use one in production.`,
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		if pkg == "" {
			pkg = "factories"
		}
		options := factoryOptions{
			pkg,
			name,
		}

		t := template.Must(template.New("factory").Parse(factoryTemplate))
		os.Mkdir(pkg, 0777)
		f, err := os.Create(fmt.Sprintf("./%s/%s.go", pkg, name))
		if err != nil {
			panic("Unable to create factory")
		}
		t.Execute(f, options)
		f.Close()
	},
}

func init() {
	makeCmd.AddCommand(factoryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// factoryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// factoryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

var factoryTemplate = `
package {{.Package}}

type {{.Name}} struct {}

func (c *{{.Name}}) Create() int {
  	return c.list.Len()
}
`
