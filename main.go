package main

import (
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
    "github.com/labstack/echo/engine/standard"
    "github.com/tuxlinuxien/gootstrap/lib/pongor"
    //"gopkg.in/flosch/pongo2.v3"
    "log"
    "flag"
    "net/http"
)

var (
    PORT string = ""
)

func init() {
    flag.StringVar(&PORT, "port", "8080", "HTTP port")
}

func home(c echo.Context) error {
    log.Println(c.Render(http.StatusOK, "pages/index.html", nil))
    return nil
}

func main() {
    flag.Parse()
    e := echo.New()
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    r := pongor.GetRenderer(pongor.PongorOption{
        //Directory: "templates",
	    Reload: true,
	})

    e.SetRenderer(r)
    e.Static("/static", "public/static")

    e.Get("/", home)

    log.Println("Server started *:", PORT)
    e.Run(standard.New(":"+PORT))
}
