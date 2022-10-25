package openapi

import (
	"fmt"
	"image"
	"os"
	"path/filepath"

	_ "github.com/hellofresh/health-go/v4/checks/postgres"
	"gorm.io/gorm"
)

func SeedMenuItems() {

	DB.Migrator().DropTable(&MenuItem{})
	DB.AutoMigrate(&MenuItem{})

	createMenuItem(DB, "Fresh from the tap", "Water", 1.99, 0, "water")
	createMenuItem(DB, "Chicken Wrap - Sandwich", "Chicken Wrap", 14.99, 1, "wrap")
	createMenuItem(DB, "A slow cooked stew", "Stew", 12.99, 2, "stew")
	createMenuItem(DB, "It looks good in the menu picture", "Tomato Soup", 4.99, 3, "soup")
	createMenuItem(DB, "A green salad", "Salad", 4.99, 4, "salad")

	var items []MenuItem
	DB.Find(&items)

	//exepath, _ := os.Executable()

	for _, item := range items {
		imagepath, _ := filepath.Abs("./go/images/" + item.ImageName + ".jpg")
		base64image := ConvertImageToBase64(imagepath)
		fmt.Println(base64image)
		item.Image = base64image
		DB.Save(&item)
	}

	DB.AutoMigrate(&MenuItem{})

}

func createMenuItem(db *gorm.DB, desc string, name string, price float32, imageid int32, imagename string) {

	err := db.Create(&MenuItem{
		Description: desc,
		Name:        name,
		Price:       price,
		ImageId:     imageid,
		ImageName:   imagename,
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
