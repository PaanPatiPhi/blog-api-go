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


    var jwks map[string]interface{}
    json.Unmarshal(body, &jwks)
    return jwks, nil
}


// แค่ verify token และ set userID — ใช้กับ route ทั่วไป
func AuthRequired(c *fiber.Ctx) error {
    authHeader := c.Get("Authorization")

    if !strings.HasPrefix(authHeader, "Bearer ") {

        return c.Status(401).JSON(fiber.Map{"error": "unauthorized"})
    }
    tokenString := strings.TrimPrefix(authHeader, "Bearer ")

    token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
        if _, ok := t.Method.(*jwt.SigningMethodECDSA); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
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

        return c.Status(401).JSON(fiber.Map{"error": "invalid token"})
    }

    claims := token.Claims.(jwt.MapClaims)
    userID := claims["sub"].(string)
    c.Locals("userID", userID)
    return c.Next()
}

// verify token + เช็คว่าเป็น admin — ใช้กับ admin routes
func AdminOnly(c *fiber.Ctx) error {
    authHeader := c.Get("Authorization")
    if !strings.HasPrefix(authHeader, "Bearer ") {
        return c.Status(401).JSON(fiber.Map{"error": "unauthorized"})
    }
    tokenString := strings.TrimPrefix(authHeader, "Bearer ")

    // parse โดยดึง public key จาก Supabase JWKS
    token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
        if _, ok := t.Method.(*jwt.SigningMethodECDSA); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
        }

        // ดึง JWKS จาก Supabase
        jwks, err := getJWKS()
        if err != nil {
            return nil, err
        }

        // parse public key จาก JWKS
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
        return c.Status(404).JSON(fiber.Map{"error": "user not found"})
    }

    if role != "admin" {
        return c.Status(403).JSON(fiber.Map{"error": "forbidden"})
    }

    c.Locals("userID", userID)
    c.Locals("role", role)
    return c.Next()
}