# ghast

[![Build Status](https://travis-ci.org/bradcypert/ghast.svg?branch=master)](https://travis-ci.org/bradcypert/ghast)

Ghast is an "All in one" toolkit for building reliable Go web applications.

Whether you're building an API, website, or something a little more exotic, Ghast has got your back. Ghast is a collection of common functionality that I've extracted, refactored, and built upon from several different Golang APIs and takes inspiration from classics such as "Rails" and "Laravel".

Here's what you should know about Ghast:

1. It's lightweight. The framework should be seen as bare helpers to the standard library.
2. We support any database that Gorm supports because Ghast uses Gorm.
3. Ghast currently follows the MVC paradigm closely. If this isn't your cup of tea, you MAY still benefit from Ghast, but it really works best when all pieces play together nicely.
4. Ghast ships with a CLI to help you generate your ghast controllers, models, and much more!

# Ghast CLI

Still here? Great! You can install Ghast's CLI by running:

```bash
go get github.com/bradcypert/ghast
```

### Create a new Ghast project

```bash
ghast new MyProject
```

### Create a new Controller

```bash
ghast make controller UsersController
```

Next steps: Open the [GhastApp README](./pkg/app/README.md) to learn how your generated code fits together!

## Learn More via Module-Specific READMEs
 - [Learn how to leverage GhastApp, the core module behind the Ghast Framework](./pkg/app/README.md).
 - [Learn how to set up the GhastRouter and respond to HTTP requests](./pkg/router/README.md).
 - [Learn how to organize your code via GhastControllers and simplify your request handlers](./pkg/controllers/README.md).