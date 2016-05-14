package main

import (
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
    "github.com/labstack/echo/engine/standard"
    "github.com/tuxlinuxien/gootstrap/lib/pongor"
    "github.com/tuxlinuxien/gootstrap/routes/account"
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

func auth() {


	//fmt.Println("Visit the URL for the auth dialog: %v", url)


    // var code string
	// if _, err := fmt.Scan(&code); err != nil {
	// 	log.Fatal(err)
	// }
	// _, err := conf.Exchange(oauth2.NoContext, code)
	// if err != nil {
	// 	log.Fatal(err)
	// }

    // tok, err := conf.Exchange(oauth2.NoContext, "authorization-code")
    // if err != nil {
    //     log.Fatal(err)
    // }
    // log.Println("token", tok, err)
    //client := conf.Client(oauth2.NoContext, tok)
    //log.Println(client.Get("/authorizations"))

}

func main() {
    //auth()
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
    account.Init(e)

    log.Println("Server started *:", PORT)
    e.Run(standard.New(":"+PORT))
}
