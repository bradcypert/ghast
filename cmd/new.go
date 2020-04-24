package cmd

import (
	"fmt"
	"os"
	"text/template"

	"github.com/spf13/cobra"
)

// newCmd cobra command to help generate a new Ghast project
var newCmd = &cobra.Command{
	Use:   "new",
	Args:  cobra.MinimumNArgs(1),
	Short: "Create a new Ghast project",
	Long:  `Creates a new Ghast project based off the provided project name.`,
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]
		t := template.Must(template.New("main").Parse(mainTemplate))
		os.Mkdir(projectName, 0777)
		f, err := os.Create(fmt.Sprintf("./%s/main.go", projectName))
		if err != nil {
			panic("Unable to create new Ghast application")
		}
		t.Execute(f, nil)
		f.Close()

		fmt.Sprintln("Successfully create a new Ghast project in ./%s", projectName)
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}

var mainTemplate = `package main

import (
	"fmt"
	"log"
	"net/http"

	ghastRouter "github.com/bradcypert/ghast/pkg/router"
)

func main() {
	router := ghastRouter.Router{}

	// Want to use controllers instead? Try running "ghast make controller MyController" from your terminal
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Hello from Ghast!")
	})
	s := router.DefaultServer()
	log.Fatal(s.ListenAndServe())
}
`
