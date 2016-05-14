package main

import (
    "github.com/gin-gonic/gin"
    "github.com/gin-gonic/contrib/sessions"
    "github.com/robvdl/pongo2gin"
    "github.com/flosch/pongo2"
    "net/http"
    "github.com/tuxlinuxien/gootstrap/routes/account"
    _ "github.com/tuxlinuxien/gootstrap/models"
    "github.com/tuxlinuxien/gootstrap/config"
)

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

    router := gin.New()
    router.Use(gin.Recovery())
    router.HTMLRender = pongo2gin.Default()

    store := sessions.NewCookieStore([]byte("secret"))
    router.Use(sessions.Sessions("mysession", store))
    router.Static("/static", "public/static")

    router.GET("/", home)
    account.Init(router)

    router.Run(":"+config.Get("port").(string))
}
