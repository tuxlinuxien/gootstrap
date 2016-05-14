package account

import (
    "github.com/labstack/echo"
    "gopkg.in/flosch/pongo2.v3"
    "net/http"
    "log"
    "golang.org/x/oauth2"
    "golang.org/x/oauth2/github"
    "io/ioutil"
    "github.com/tuxlinuxien/gootstrap/models"
    "crypto/sha1"
    "fmt"
    "time"
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

func Init(e *echo.Echo) {
    accountGroup := e.Group("/account")
    {
        accountGroup.Get("/login", homeGet)
        accountGroup.Post("/login", homePost)
        accountGroup.Get("/login/github", logingithub)
        accountGroup.Get("/callback", cb)
        accountGroup.Get("/register", registerGet)
        accountGroup.Post("/register", registerPost)
    }
}

func homeGet(c echo.Context) error {
    return c.Render(http.StatusOK, "account/login.html", nil)
}

func homePost(c echo.Context) error {
    var user_post struct {
        Email string `form:"email"`
        Pwd string `form:"password"`
    }
    if err := c.Bind(&user_post); err != nil {
        return err
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
        return c.Render(http.StatusOK, "account/login.html", pongo2.Context{
            "error": "User not found.",
        })
    }
    createCookie(c, u)
    return c.Redirect(302, "/")
}

func registerGet(c echo.Context) error {
    return c.Render(http.StatusOK, "account/register.html", nil)
}

func registerPost(c echo.Context) error {
    var user_post struct {
        Email string `form:"email"`
        Pwd string `form:"password"`
        PwdC string `form:"password_repeat"`
    }
    if err := c.Bind(&user_post); err != nil {
        return err
    }
    u := new(models.User)
    u.Email = user_post.Email
    u.Password = fmt.Sprintf("%x", sha1.Sum([]byte(user_post.Pwd)))
    _, err := models.Engine.Insert(u)
    if err != nil {
        return c.Render(http.StatusOK, "account/register.html", pongo2.Context{
            "error": "User already exists.",
        })
    }
    return c.Render(http.StatusOK, "account/register.html", pongo2.Context{
        "ok": "User created",
    })
}

func logingithub(c echo.Context) error {
    url := CONF.AuthCodeURL("state", oauth2.AccessTypeOffline)
    return c.Redirect(302, url)
}

func cb(c echo.Context) error {
    code := c.QueryParam("code")
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
    return c.Redirect(302, "/")
}

func createCookie(c echo.Context, u *models.User) {
    cookie := new(echo.Cookie)
	cookie.SetName("email")
	cookie.SetValue(u.Email)
    cookie.SetSecure(true)
    cookie.SetExpires(time.Now().Add(2400 * time.Hour))
    c.SetCookie(cookie)
}
