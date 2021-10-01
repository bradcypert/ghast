package cmd

import (
	// nolint
	_ "embed"
	"fmt"
	"os"
	"text/template"

	"github.com/spf13/cobra"
)

type modelOptions struct {
	Package string
	Name    string
}

// modelCmd represents the model command
var modelCmd = &cobra.Command{
	Use:   "model",
	Short: "create a new model",
	Long: `Create a new model for your ghast project.
    
    Models leverage GORM and are used to structure the data that you're
    working with in a way that makes sense to your application's needs.`,
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		if pkg == "" {
			pkg = "models"
		}
		options := modelOptions{
			pkg,
			name,
		}

		t := template.Must(template.New("model").Parse(modelTemplate))
		os.Mkdir(pkg, 0777)
		f, err := os.Create(fmt.Sprintf("./%s/%s.go", pkg, name))
		if err != nil {
			panic("Unable to create model")
		}
		t.Execute(f, options)
		f.Close()
	},
}

func init() {
	makeCmd.AddCommand(modelCmd)
}

//go:embed model.tmpl
var modelTemplate string
