package account

import (
    "github.com/gin-gonic/gin"
    "github.com/gin-gonic/contrib/sessions"
    "github.com/flosch/pongo2"
    "net/http"
    "log"
    "golang.org/x/oauth2"
    "golang.org/x/oauth2/github"
    "io/ioutil"
    "github.com/tuxlinuxien/gootstrap/models"
    "crypto/sha1"
    "fmt"
    "encoding/json"
)

var (
    CONF *oauth2.Config = nil
)

func init() {
    CONF = &oauth2.Config{
        ClientID:     "78b1337a2bd250a39318",
        ClientSecret: "8a612f8d59dedd67436d7719b18d3d822968b601",
        RedirectURL:  "http://localhost:8080/account/callback",
        Scopes: []string{
            "user:email",
        },
        Endpoint: github.Endpoint,
    }
}

func Init(e *gin.Engine) {
    accountGroup := e.Group("/account")
    {
        accountGroup.GET("/login", homeGet)
        accountGroup.POST("/login", homePost)
        accountGroup.GET("/login/github", logingithub)
        accountGroup.GET("/callback", cb)
        accountGroup.GET("/register", registerGet)
        accountGroup.POST("/register", registerPost)
        accountGroup.GET("/user", userPage)
        accountGroup.GET("/logout", logout)
    }
}

func homeGet(c *gin.Context) {
    c.HTML(http.StatusOK, "account/login.html", pongo2.Context{})
}

func homePost(c *gin.Context) {
    var user_post struct {
        Email string `form:"email"`
        Pwd string `form:"password"`
    }
    if err := c.Bind(&user_post); err != nil {
        log.Fatal("Error login")
    }
    user_post.Pwd = fmt.Sprintf("%x", sha1.Sum([]byte(user_post.Pwd)))

    u := &models.User{}
    has, err := models.Engine.
    Where("Email = ?", user_post.Email).
    And("Password = ?", user_post.Pwd).
    Get(u)
    if err != nil {
        log.Fatal("Error login")
    }
    if has != true {
        c.HTML(http.StatusOK, "account/login.html", pongo2.Context{
            "error": "User not found.",
        })
        return
    }
    session := sessions.Default(c)
    session.Set("email", u.Email)
    session.Save()
    c.Redirect(302, "/")
}

func registerGet(c *gin.Context) {
    c.HTML(http.StatusOK, "account/register.html", pongo2.Context{})
}

func registerPost(c *gin.Context) {
    var user_post struct {
        Email string `form:"email"`
        Pwd string `form:"password"`
        PwdC string `form:"password_repeat"`
    }
    if err := c.Bind(&user_post); err != nil {
        c.String(500, "error")
        return
    }
    u := new(models.User)
    u.Email = user_post.Email
    u.Password = fmt.Sprintf("%x", sha1.Sum([]byte(user_post.Pwd)))
    u.Type = "db"
    _, err := models.Engine.Insert(u)
    if err != nil {
        c.HTML(http.StatusOK, "account/register.html", pongo2.Context{
            "error": "User already exists.",
        })
        return
    }
    c.HTML(http.StatusOK, "account/register.html", pongo2.Context{
        "ok": "User created",
    })
    return
}

func logingithub(c *gin.Context) {
    url := CONF.AuthCodeURL("state", oauth2.AccessTypeOffline)
    c.Redirect(302, url)
}

func cb(c *gin.Context) {
    code := c.Query("code")
    tok, err := CONF.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Fatal(err)
	}
    client := CONF.Client(oauth2.NoContext, tok)
	resp, err := client.Get("https://api.github.com/user")
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    log.Println(string(body), err)
    var r struct {
        Email string `json:"email"`
    }
    json.Unmarshal(body, &r)
    u := &models.User{}
    has, _ := models.Engine.
    Where("Email = ?", r.Email).
    Get(u)
    if has == false {
        u := new(models.User)
        u.Email = r.Email
        u.Type = "github"
        models.Engine.Insert(u)
    }
    session := sessions.Default(c)
    session.Set("email", r.Email)
    session.Save()
    c.Redirect(302, "/")
}

func userPage(c *gin.Context) {
    session := sessions.Default(c)
    v := session.Get("email")
    if v == nil {
        c.Redirect(302, "/account/login")
        return
    }
    u := &models.User{}
    models.Engine.
    Where("Email = ?", v.(string)).
    Get(u)
    c.HTML(http.StatusOK, "account/user.html", pongo2.Context{
        "user": u,
    })
}

func logout(c *gin.Context) {
    session := sessions.Default(c)
    session.Clear()
    session.Save()
    c.Redirect(302, "/")
}
