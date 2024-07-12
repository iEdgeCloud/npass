package main

import (
    "github.com/kataras/iris/v12"
    "github.com/kataras/iris/v12/middleware/jwt"
    "github.com/iedgecloud/npass/handlers"
    "github.com/iedgecloud/npass/database"
)

var jwtHandler *jwt.Middleware

func main() {
    app := iris.New()

    // Set logger to debug level
    app.Logger().SetLevel("debug")

    // Log all requests
    app.Use(iris.Logger())

    database.InitDB()

    // Initialize JWT middleware
    jwtHandler = jwt.New(jwt.Config{
        SigningMethod: "HS256",
        SigningKey:    []byte("secret_key"),
        Expiration:    true,
        TokenLookup:   "header:Authorization",
        Extractor:     jwt.FromAuthHeader,
    })

    app.Post("/api/login", handlers.Login)
    app.Get("/api/config", jwtHandler.Serve, handlers.GetConfig)

    // Start the server
    app.Listen(":8080", iris.WithConfiguration(iris.Configuration{
        DisableStartupLog: false,
        EnableOptimizations: true,
    }))
}

