package cmd

import (
	"fmt"
	"os"
	"text/template"

	"github.com/spf13/cobra"
)

// migrationCmd represents the migration command
var migrationCmd = &cobra.Command{
	Use:   "migration",
	Short: "Create a new database migration",
	Long: `Create a new database migration.
    
    Database migrations help manage your application's schema
    and it's changes over time. Migrations are written in Go and use GORM (ideally).
    `,
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		if pkg == "" {
			pkg = "migrations"
		}
		options := factoryOptions{
			pkg,
			name,
		}

		t := template.Must(template.New("migration").Parse(migrationTemplate))
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
	makeCmd.AddCommand(migrationCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// migrationCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// migrationCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

var migrationTemplate = `
// TODO
`
