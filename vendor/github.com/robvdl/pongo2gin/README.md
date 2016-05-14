Pongo2gin
=========

Package pongo2gin is a template renderer that can be used with the Gin web
framework https://github.com/gin-gonic/gin it uses the Pongo2 template library
https://github.com/flosch/pongo2

Usage
-----

To use pongo2gin you need to set your router.HTMLRenderer to a new renderer
instance, this is done after creating the Gin router when the Gin application
starts up. You can use pongo2gin.Default() to create a new renderer with
default options, this assumes templates will be located in the "templates"
directory, or you can use pongo2.New() to specify a custom location.

To render templates from a route, call c.HTML just as you would with
regular Gin templates, the only difference is that you pass template
data as a pongo2.Context instead of gin.H type.

Basic Example
-------------

```go
import (
    "github.com/gin-gonic/gin"
    "github.com/flosch/pongo2"
    "github.com/robvdl/pongo2gin"
)

func main() {
    router := gin.Default()

    // Use pongo2gin.Default() for default options or pongo2gin.New()
    // if you need to use custom RenderOptions.
    router.HTMLRender = pongo2gin.Default()

    router.GET("/", func(c *gin.Context) {
        // Use pongo2.Context instead of gin.H
        c.HTML(200, "hello.html", pongo2.Context{"name": "world"})
    })

    router.Run(":8080")
}
```

RenderOptions
-------------

When calling pongo2gin.New() instead of pongo2gin.Default() you can use these
custom RenderOptions:

```go
type RenderOptions struct {
    TemplateDir string  // location of the template directory
    ContentType string  // Content-Type header used when calling c.HTML()
}
```

Template Caching
----------------

Templates will be cached if the current Gin Mode is set to anything but "debug",
this means the first time a template is used it will still load from disk, but
after that the cached template will be used from memory instead.

If he Gin Mode is set to "debug" then templates will be loaded from disk on
each request.

Caching is implemented by the Pongo2 library itself.

GoDoc
-----

https://godoc.org/github.com/robvdl/pongo2gin
