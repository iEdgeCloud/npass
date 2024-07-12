package handlers

import (
    "github.com/kataras/iris/v12"
    "github.com/kataras/iris/v12/middleware/jwt"
    "github.com/iedgecloud/npass/models"
    "time"
)

type LoginRequest struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

func Login(ctx iris.Context) {
    var req LoginRequest
    if err := ctx.ReadJSON(&req); err != nil {
        ctx.StatusCode(iris.StatusBadRequest)
        ctx.JSON(iris.Map{"error": err.Error()})
        return
    }

    user, err := models.Authenticate(req.Username, req.Password)
    if err != nil {
        ctx.StatusCode(iris.StatusUnauthorized)
        ctx.JSON(iris.Map{"error": "Invalid credentials"})
        return
    }

    token, err := generateJWT(user.ID)
    if err != nil {
        ctx.StatusCode(iris.StatusInternalServerError)
        ctx.JSON(iris.Map{"error": "Error generating token"})
        return
    }

    ctx.JSON(iris.Map{"token": token})
}

func generateJWT(userID int) (string, error) {
    token := jwt.NewTokenWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(time.Hour * 24).Unix(),
    })

    tokenString, err := token.SignedString([]byte("secret_key"))
    if err != nil {
        return "", err
    }

    return tokenString, nil
}

