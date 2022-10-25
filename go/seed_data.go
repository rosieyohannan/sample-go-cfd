package openapi

import (
	"image"
	"os"
	"path/filepath"

	_ "github.com/hellofresh/health-go/v4/checks/postgres"
	"gorm.io/gorm"
)

func SeedMenuItems() {

	DB.Migrator().DropTable(&MenuItem{})
	DB.AutoMigrate(&MenuItem{})

	createMenuItem(DB, "Fresh from the tap", "Water", 1.99, 0)
	createMenuItem(DB, "Chicken Wrap - Sandwich", "Chicken Wrap", 14.99, 1)
	createMenuItem(DB, "A slow cooked stew", "Stew", 12.99, 2)
	createMenuItem(DB, "It looks good in the menu picture", "Tomato Soup", 4.99, 3)
	createMenuItem(DB, "A green salad", "Salad", 4.99, 4)
}

func createMenuItem(db *gorm.DB, desc string, name string, price float32, imageid int32) {

	err := db.Create(&MenuItem{
		Description: desc,
		Name:        name,
		Price:       price,
		ImageId:     imageid,
	}).Error
	if err != nil {
		panic(err)
	}
}

func getImageFromFilePath(filePath string) (image.Image, error) {
	path := filepath.Join(".\\images", filePath)
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	image, _, err := image.Decode(f)
	return image, err
}
