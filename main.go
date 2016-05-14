package main

import (
    "github.com/gin-gonic/gin"
    "github.com/gin-gonic/contrib/sessions"
    "github.com/robvdl/pongo2gin"
    "github.com/flosch/pongo2"
    "log"
    "flag"
    "net/http"
    "github.com/tuxlinuxien/gootstrap/routes/account"
    _ "github.com/tuxlinuxien/gootstrap/models"
)

var (
    PORT string = ""
)

func init() {
    flag.StringVar(&PORT, "port", "8080", "HTTP port")
}

func home(c *gin.Context) {
    session := sessions.Default(c)
    v := session.Get("email")
    if v != nil {
        c.Redirect(302, "/account/user")
        return
    }
    c.HTML(http.StatusOK, "pages/index.html", pongo2.Context{})
}

func main() {
    flag.Parse()

    router := gin.New()
    router.Use(gin.Recovery())
    router.HTMLRender = pongo2gin.Default()

    store := sessions.NewCookieStore([]byte("secret"))
    router.Use(sessions.Sessions("mysession", store))
    router.Static("/static", "public/static")

    router.GET("/", home)
    account.Init(router)

    log.Println("Server started *:", PORT)
    router.Run(":"+PORT)
}
