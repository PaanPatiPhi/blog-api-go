package middlewares

import (
	"blog-api-go/database"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"fmt"
	"github.com/go-jose/go-jose/v3"
)


func Logger(c *fiber.Ctx) error {
    // log ทุก request
    err := c.Next()
    log.Printf("%s %s → %d", c.Method(), c.Path(), c.Response().StatusCode())
    return err
}

func getJWKS() (map[string]interface{}, error) {
    url := os.Getenv("SUPABASE_URL") + "/auth/v1/.well-known/jwks.json"
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    body, _ := io.ReadAll(resp.Body)

     log.Println("JWKS response:", string(body)) 

    var jwks map[string]interface{}
    json.Unmarshal(body, &jwks)
    return jwks, nil
}

func AdminOnly(c *fiber.Ctx) error {
    authHeader := c.Get("Authorization")
    if !strings.HasPrefix(authHeader, "Bearer ") {
        return c.Status(401).JSON(fiber.Map{"error": "unauthorized"})
    }
    tokenString := strings.TrimPrefix(authHeader, "Bearer ")
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        jwks, err := getJWKS()
        if err != nil{
            return nil, err
        }
        keys := jwks["keys"].([]interface{})
        keyData, _ := json.Marshal(keys[0])

        var ecKey jose.JSONWebKey
        json.Unmarshal(keyData, &ecKey)
        return ecKey.Key, nil

    })
    if err != nil || !token.Valid {
        return c.Status(401).JSON(fiber.Map{"error": "invalid token", "details": err.Error()})
    }

    claims := token.Claims.(jwt.MapClaims)
    userID := claims["sub"].(string)
    
    var role string
    err = database.RawDB.Get(&role, "SELECT role FROM users WHERE id = $1", userID)
    if err != nil {
        return c.Status(401).JSON(fiber.Map{"error": "invalid token, cannot get role", "details": err.Error()})
    }
    if role != "admin" {
        return c.Status(403).JSON(fiber.Map{"error": "forbidden"})
    }
    
    c.Locals("userID", userID)
    c.Locals("role", role)
    return c.Next()
}
