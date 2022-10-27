package openapi

import (
	"encoding/base64"
	"errors"
	"fmt"
	"image"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func ConvertImageToBase64(filepath string) string {
	// Read the entire file into a byte slice
	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	var base64Encoding string

	// Determine the content type of the image file
	mimeType := http.DetectContentType(bytes)

	// Prepend the appropriate URI scheme header depending
	// on the MIME type
	switch mimeType {
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	}

	// Append the base64 encoded output
	base64Encoding += toBase64(bytes)

	// return the string representation of the image
	return base64Encoding
}

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func decodeJpgImage(path string) (image.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fi, _ := f.Stat()
	fmt.Println(fi.Name())
	//defer f.Close()sss
	img, format, err := image.Decode(f)
	if err != nil {
		fmt.Println("Decoding error:", err.Error())
		return nil, err
	}
	if format != "jpeg" {
		fmt.Println("image format is not jpeg")
		return nil, errors.New("")
	}
	return img, nil
}
