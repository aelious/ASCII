package ASCII

import (
	"fmt"
	"image"
	_ "image/png" // import this package to decode PNGs
	"os"
)

func GetAndConvertColors() {
	fmt.Printf("Processing pixel RGB.\nOpening image. . .\n")
	reader, err := os.Open("gray_image.png")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	} else {
		fmt.Print("Image opened successfully.\n")
	}
	defer reader.Close()

	img, _, err := image.Decode(reader)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	} else {
		fmt.Print("Image decoded successfully.\n")
	}

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	fmt.Printf("Creating output file for pixels.\n")
	// Create outfile
	fo, err := os.Create("pixels_RGB.txt")
	// Close outfile on exit
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()
	asciiImg, err := os.Create("asciiImg.txt")

	fmt.Printf("File successfully created.\n")
	fmt.Printf("Writing pixels to file... This may take some time.\n")
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			color := img.At(x, y)
			r, g, b, _ := color.RGBA()
			brightness := (0.299 * float32(r>>8)) + (0.587 * float32(g>>8)) + (0.116 * float32(b>>8))
			_, err := asciiImg.WriteString(fmt.Sprintf("%f\n", brightness))
			if err != nil {
				panic(err)
			}
			_, errors := fo.WriteString(fmt.Sprintf("Pixel at (%d, %d) - R: %d, G: %d, B: %d\n", x, y, r>>8, g>>8, b>>8))
			if errors != nil {
				panic(errors)
			}
		}
	}
	fmt.Printf("File saved successfully.\n")
}
