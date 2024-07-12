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

    app.Use(iris.Logger())


    database.InitDB()  // 初始化数据库

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

    app.Listen(":8080")
}

