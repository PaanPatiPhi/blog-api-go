package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func AdminOnly(c *fiber.Ctx) error {
    // TODO: ดึง role จาก JWT claims ทีหลัง
    // ตอนนี้ placeholder ไว้ก่อน
    role := c.Get("X-Role")
    if role != "admin" {
        return c.Status(403).JSON(fiber.Map{"error": "forbidden"})
    }
    return c.Next()
}

func Logger(c *fiber.Ctx) error {
    // log ทุก request
    err := c.Next()
    log.Printf("%s %s → %d", c.Method(), c.Path(), c.Response().StatusCode())
    return err
}