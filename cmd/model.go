package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var pkg string
var name string

// modelCmd represents the model command
var modelCmd = &cobra.Command{
	Use:   "model",
	Short: "create a new model",
	Long: `Create a new model for your ghast project.
	
	Models leverage GORM and are used to structure the data that you're
	working with in a way that makes sense to your application's needs.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("model called")
	},
}

func init() {
	makeCmd.AddCommand(modelCmd)
	controllerCmd.Flags().StringVarP(&pkg, "package", "p", "factories", "Package name")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// modelCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// modelCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

var factoryTemplate = `
package {{.Package}}

import (
	"github.com/jinzhu/gorm"
)

type {{.Name}} struct {
	gorm.Model
}
`
