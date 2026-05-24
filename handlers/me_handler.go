package handlers

import (
	"blog-api-go/repositories"
	"database/sql"
	"os"

	storage_go "github.com/supabase-community/storage-go"
	"github.com/gofiber/fiber/v2"
)

func GetMeProfile(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	me, err := repositories.GetMe(userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(404).JSON(fiber.Map{"error": "user not found"})
		}
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to get user",
		})
	}
	return c.JSON(me)
}

func UpdateUserProfile(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	username := c.FormValue("username")
	name := c.FormValue("name")
	profilePic := ""

	file, err := c.FormFile("imageFile")
	if err == nil {
		f, err := file.Open()
		if err != nil{
			return c.Status(500).JSON(fiber.Map{
				"error": "Failed to open file",
			})
		}
		defer f.Close()

		storageClient := storage_go.NewClient(
			os.Getenv("SUPABASE_URL")+"/storage/v1",
			os.Getenv("SUPABASE_SERVICE_ROLE_KEY"),
			nil,
		)

		filePath := "profile/" + userID + "_" + file.Filename
		contentType := file.Header.Get("Content-Type")

		_, err = storageClient.UploadFile("pics", filePath, f, storage_go.FileOptions{
			ContentType: &contentType,
		})
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Failed to upload file",
			})
		}
		profilePic = os.Getenv("SUPABASE_URL") + "/storage/v1/object/public/pics/" + filePath
	}
	var usernamePtr, namePtr, profilePicPtr *string
	if username != "" {
		usernamePtr = &username
	}
	if name != "" {
		namePtr = &name
	}
	if profilePic != "" {
		profilePicPtr = &profilePic
}

if err := repositories.UpdateUserProfile(userID, usernamePtr, namePtr, profilePicPtr); err != nil {
    return c.Status(500).JSON(fiber.Map{"error": "failed to update profile"})
}
return c.JSON(fiber.Map{
    "message": "User profile updated successfully",
})
}