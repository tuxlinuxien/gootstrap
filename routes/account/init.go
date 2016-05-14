package account

import (
    "github.com/labstack/echo"
    //"gopkg.in/flosch/pongo2.v3"
    "net/http"
    "log"
    "golang.org/x/oauth2"
    "golang.org/x/oauth2/github"
    "io/ioutil"
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
        accountGroup.Get("/callback", cb)
    }
}

func homeGet(c echo.Context) error {
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
    return c.Render(http.StatusOK, "account/login.html", nil)
}
