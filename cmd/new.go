package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
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
		runDir, err := os.Getwd()
		if err != nil {
			log.Fatal("Unable to get working directory when creating new ghast app")
		}
		fmt.Print("Please enter your root package name: ")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		pkgName := strings.Replace(text, "\n", "", -1)

		// make relevant directories
		os.Mkdir(projectName, 0777)
		os.Mkdir(projectName+"/views", 0777)
		os.Mkdir(projectName+"/controllers", 0777)

		type pkg struct {
			Pkg string
		}
		// make mod file
		modFileTemplate := template.Must(template.New("mod").Parse(modTemplate))
		os.Chdir("./" + projectName)
		f, err := os.Create("./go.mod")
		if err != nil {
			panic("Unable to create new Ghast application controller")
		}
		modFileTemplate.Execute(f, pkg{pkgName})
		f.Close()
		os.Chdir(runDir)

		// make initial controller
		controllerTemplate := template.Must(template.New("controller").Parse(demoControllerTemplate))
		os.Chdir("./" + projectName + "/controllers")
		f, err = os.Create("./HomeController.go")
		if err != nil {
			panic("Unable to create new Ghast application controller")
		}
		controllerTemplate.Execute(f, nil)
		f.Close()
		os.Chdir(runDir)

		// make initial view
		viewTemplate := template.Must(template.New("view").Parse(viewTemplate))
		os.Chdir("./" + projectName + "/views")
		f, err = os.Create("./template.jet")
		if err != nil {
			panic("Unable to create new Ghast application template")
		}
		viewTemplate.Execute(f, nil)
		f.Close()
		os.Chdir(runDir)

		// make main file
		mainTemplate := template.Must(template.New("main").Parse(mainTemplate))
		f, err = os.Create(fmt.Sprintf("./%s/main.go", projectName))
		if err != nil {
			panic("Unable to create new Ghast application")
		}
		mainTemplate.Execute(f, nil)
		f.Close()

		os.Chdir("./" + projectName)
		// finally fetch the go modules we need for ghast.
		goExecutable, err := exec.LookPath("go")
		cmdGoGet := &exec.Cmd{
			Path:   goExecutable,
			Args:   []string{goExecutable, "get", "-u", "./..."},
			Stdout: os.Stdout,
			Stdin:  os.Stdin,
		}

		cmdGoGet.Run()
		// if err = cmdGoGet.Run(); err != nil {
		// 	panic("Unable to fetch go modules")
		// }

		fmt.Printf("Successfully create a new Ghast project in ./%s", projectName)
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}

var modTemplate = `module {{.Pkg}}

go 1.13
`

var demoControllerTemplate = `
package controllers

import (
	"net/http"
	ghastController "github.com/bradcypert/ghast/pkg/controllers"
)

type HomeController struct {
	ghastController.GhastController
}

func (c HomeController) Index(w http.ResponseWriter, r *http.Request) {
  	c.View("template.jet", w, nil, nil)
}
`

var viewTemplate = `
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <title>Hello from Ghast!</title>
    <link rel="stylesheet" href="style.css">
    <script src="script.js"></script>
  </head>
  <body>
    <h1>Hello from Ghast!</h1>
  </body>
</html>
`

var mainTemplate = `package main

import (
	"fmt"
	"log"
	"net/http"

	ghastApp "github.com/bradcypert/ghast/pkg/app"
	ghastRouter "github.com/bradcypert/ghast/pkg/router"
)

func main() {
	router := ghastRouter.Router{}

	// Want to use controllers instead? Try running "ghast make controller MyController" from your terminal
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Hello from Ghast!")
	})

	router.Get("/:name", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Hello "+r.Context().Value("name").(string))
	})
	
	app := ghastApp.NewApp()
	app.SetRouter(router)
	app.Start()
}
`
