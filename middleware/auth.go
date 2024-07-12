package middleware

import (
    "myapp/utils"
    "github.com/kataras/iris/v12"
)

func AuthMiddleware(ctx iris.Context) {
    token := ctx.GetHeader("Authorization")
    if token == "" {
        ctx.StatusCode(iris.StatusForbidden)
        ctx.JSON(iris.Map{"error": "Forbidden"})
        return
    }

    userID, err := utils.ValidateJWT(token)
    if err != nil {
        ctx.StatusCode(iris.StatusForbidden)
        ctx.JSON(iris.Map{"error": "Forbidden"})
        return
    }

    ctx.Values().Set("user_id", userID)
    ctx.Next()
}

