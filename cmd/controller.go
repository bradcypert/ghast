package cmd

import (
	// nolint
	_ "embed"
	"fmt"
	"os"
	"text/template"

	"github.com/spf13/cobra"
)

type controllerOptions struct {
	Package string
	Name    string
}

// controllerCmd represents the controller command
var controllerCmd = &cobra.Command{
	Use:   "controller",
	Args:  cobra.MinimumNArgs(1),
	Short: "Create a controller",
	Long: `Create a controller

Create a new controller. Controllers are used to handle your application specific logic
and are delegated to by your router.`,
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		if pkg == "" {
			pkg = "controllers"
		}
		options := controllerOptions{
			pkg,
			name,
		}

		t := template.Must(template.New("controller").Parse(controllerTemplate))
		os.Mkdir(pkg, 0777)
		f, err := os.Create(fmt.Sprintf("./%s/%s.go", pkg, name))
		if err != nil {
			panic("Unable to create controller")
		}
		t.Execute(f, options)
		f.Close()
		fmt.Printf("created new controller: ./%s/%s.go\n", pkg, name)
	},
}

func init() {
	makeCmd.AddCommand(controllerCmd)
}

//go:embed controller.tmpl
var controllerTemplate string
