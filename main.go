package main

import (
	"bufio"
	"fmt"
	"image"
	"log"
	"math"
	"os"
	"path/filepath"
	"strings"

	"image/color"
	"image/draw"
	"image/gif"
	_ "image/gif"
	"image/jpeg"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
)

func main() {
	borderPadding := 300                        // 300 pixels
	borderColor := color.RGBA{255, 255, 255, 0} // White

	if len(os.Args) < 2 {
		log.Fatal("Usage: photosquared <image>\n")
	}
	inputImageFullPath := os.Args[1]
	fmt.Printf("inputImage: [%s]\n", inputImageFullPath)

	inputImageFile, err := os.Open(inputImageFullPath)
	if err != nil {
		log.Fatal(err)
	}
	defer inputImageFile.Close()

	inputImage, inputImageType, err := image.Decode(bufio.NewReader(inputImageFile))
	if err != nil {
		log.Fatalf("err: %s", err)
	}

	inputWidth := inputImage.Bounds().Max.X
	inputHeight := inputImage.Bounds().Max.Y

	outputWidthHeight := int(math.Max(float64(inputWidth), float64(inputHeight))) + (2 * borderPadding)

	fmt.Printf("Width: %d; Height: %d; outputWidthHeight: %d\n", inputWidth, inputHeight, outputWidthHeight)

	// Compute the upper left,right coordinates to place the input Image
	//   within the output Image
	var upperLeftPoint image.Point
	if inputWidth > inputHeight {
		// landscape
		yOffset := (outputWidthHeight - inputHeight) / 2
		upperLeftPoint = image.Point{borderPadding, yOffset}
	} else if inputHeight > inputWidth {
		// portrait
		xOffset := (outputWidthHeight - inputWidth) / 2
		upperLeftPoint = image.Point{xOffset, borderPadding}
	} else {
		// square
		upperLeftPoint = image.Point{borderPadding, borderPadding}
	}

	// Open new blank image
	inputImageFileDir, inputImageFilename := filepath.Split(inputImageFullPath)
	inputImageFileExt := filepath.Ext(inputImageFilename)
	outputImageFilename := fmt.Sprintf(
		"%s%s-square%s",
		inputImageFileDir,
		strings.TrimSuffix(inputImageFilename, inputImageFileExt),
		inputImageFileExt,
	)
	fmt.Printf("Writing to: [%s]\n", outputImageFilename)

	// Open new Image File
	outputImageFile, err := os.Create(outputImageFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer outputImageFile.Close()

	outputImage := image.NewRGBA(image.Rect(0, 0, outputWidthHeight, outputWidthHeight))
	draw.Draw(outputImage, outputImage.Bounds(), &image.Uniform{borderColor}, image.ZP, draw.Src)
	draw.Draw(
		outputImage,
		image.Rectangle{upperLeftPoint, outputImage.Bounds().Max},
		inputImage,
		image.Point{0, 0},
		draw.Over,
	)

	switch inputImageType {
	case "gif":
		gif.Encode(outputImageFile, outputImage, nil)
	case "jpeg":
		jpeg.Encode(outputImageFile, outputImage, nil)
	case "png":
		png.Encode(outputImageFile, outputImage)
	default:
		log.Fatalf("Unknown image type: [%s]", inputImageType)
	}
}
