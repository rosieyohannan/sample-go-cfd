package openapi

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"path/filepath"

	_ "github.com/hellofresh/health-go/v4/checks/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSLMode  string
}

func (Cfg PostgresConfig) String() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		Cfg.Host, Cfg.Port, Cfg.User, Cfg.Password, Cfg.Database, Cfg.SSLMode)

}

func OpenDb() error {

	cfg := PostgresConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "postgres",
		Database: "cfd",
		SSLMode:  "disable",
	}

	var err error
	DB, err = gorm.Open(postgres.Open(cfg.String()), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		return err
	}

	return nil
}

func SeedMenuItems() {

	DB.Migrator().DropTable(&MenuItem{})
	DB.Migrator().DropTable(&CartItem{})
	DB.AutoMigrate(&MenuItem{})

	createMenuItem(DB, "Fresh from the tap", "Water", 1.99, 0, "water")
	createMenuItem(DB, "Chicken Wrap - Sandwich", "Chicken Wrap", 14.99, 1, "wrap")
	createMenuItem(DB, "A slow cooked stew", "Stew", 12.99, 2, "stew")
	createMenuItem(DB, "It looks good in the menu picture", "Tomato Soup", 4.99, 3, "soup")
	createMenuItem(DB, "A green salad", "Salad", 4.99, 4, "salad")

	var items []MenuItem
	DB.Find(&items)

	for _, item := range items {
		imagepath, _ := filepath.Abs("./go/images/" + item.ImageName + ".jpg")
		img, _ := decodeJpgImage(imagepath)
		buf := new(bytes.Buffer)
		jpeg.Encode(buf, img, nil)
		item.Image = buf.Bytes()
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
