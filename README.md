# ghast

[![Build Status](https://travis-ci.org/bradcypert/ghast.svg?branch=master)](https://travis-ci.org/bradcypert/ghast)

Ghast is an "All in one" toolkit for building reliable Go web applications.

Whether you're building an API, website, or something a little more exotic (still over HTTP/HTTPS protocols)
Ghast has got your back. Ghast is a collection of common functionality that I've extracted and refactored from several different Golang APIs and takes inspiration from classics such as "Rails" and "Laravel".

Here's what you should know about Ghast:

1. It's lightweight. The framework should be seen as bare helpers to the standard library.
2. We support any database that Gorm supports because Ghast uses Gorm.
3. Ghast currently follows the MVC paradigm closely. If this isn't your cup of tea, you MAY still benefit from Ghast, but it really works best when all pieces play together nicely.
4. Ghast ships with a CLI to help you generate your ghast controllers, models, and much more!

Still here? Great! You can install Ghast's CLI by running:

```
go get github.com/bradcypert/ghast
```

More documentation coming soon!
