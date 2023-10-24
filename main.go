package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
)


func readImageFile(filePath string) image.Image {
	reader, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	defer reader.Close()
	fmt.Println("Image file opened successfully")
	img, _, err := image.Decode(reader)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	fmt.Println("Image file decoded successfully")
	return img
}

func searchtExt(img image.Image) string {
	_, isPng := img.(*image.NRGBA)
	if !isPng {
		fmt.Println("Image file is not PNG")
		return ""
	}

	bounds := img.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			if r == 0x74 && g == 0x45 && b == 0x58 && a == 0xFF {
				textData := ""
				for y := y ; y < bounds.Max.Y; y++ {
					for x := bounds.Min.X; x < bounds.Max.X; x++ {
					r, g, b, a := img.At(x, y).RGBA()
					if r == 0x00 && g == 0x00 && b == 0x00 && a == 0xFF {
						return textData
					}
					textData += string(r)
					textData += string(g)
					textData += string(b)
				}
				}
				return textData
			}
		}
}
	return ""
}

func main() {
	imageFilePath := flag.String("image", "", "Image file path")
	flag.Parse()

	if *imageFilePath == "" {
		fmt.Println("Please specify image file path")
		os.Exit(1)
	}

	fmt.Println("Image file path: ", *imageFilePath)
	img := readImageFile(*imageFilePath)
	textData := searchtExt(img)
	fmt.Println("Text data: ", textData)
}
