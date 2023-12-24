package main

import (
	//"github.com/labstack/echo-jwt"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"yapper.com/m/yapper/cmd/auth"
	"yapper.com/m/yapper/cmd/handlers"
	"yapper.com/m/yapper/cmd/templates"
)

func main(){
    e := echo.New()
    e.Static("/", "static")

    e.Use(middleware.Logger())
    //We need to customize the JWT mw to redirect the user to the login page 
    // if the token is expired or simply invalid
    //e.Use(echojwt.JWT([]byte("secret")))

    templs := templates.NewTemplate()
    templs.AddTemplate("index")
    templs.AddTemplate("signup")
    templs.AddTemplate("signin")
    templs.AddTemplate("profile")
    templs.AddTemplate("edit_profile")
    templs.AddTemplate("topbar")
    e.Renderer = templs
    
    e.GET("/", handlers.Home)
    e.GET("/user/:id", handlers.UserProfile)
    e.GET("/register", handlers.GotoRegister)
    e.POST("/register", handlers.Register)
    e.GET("/signin", handlers.GotoSignin)
    e.POST("/signin", handlers.Signin)

    edit := e.Group("edit")
    config := echojwt.Config{
        NewClaimsFunc: func(c echo.Context) jwt.Claims{
            return new(auth.CustomClaim)
        },
        SigningKey: []byte("secret"),
        TokenLookup: "cookie:user",
        ErrorHandler: func(c echo.Context, err error) error {
            println(err.Error())
            return err
        },
    }
    edit.Use(echojwt.WithConfig(config))

    edit.GET("/:id", handlers.EditProfile)
    edit.POST("/:id", handlers.SaveProfileChanges)
    e.Logger.Fatal(e.Start(":1323"))
}
