package ASCII

import (
	"fmt"
	"os"

	"github.com/fogleman/gg"
	"golang.org/x/image/font/gofont/goregular"
)

func CreatePicText(s string) {
	fmt.Printf("Processing image text.\nOpening image. . .\n")
	// Create a temporary file and write the byte slice to it
	tempFile, err := os.CreateTemp("", "font-*.ttf")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tempFile.Name())

	if _, err := tempFile.Write(goregular.TTF); err != nil {
		panic(err)
	}

	dc, err := gg.LoadPNG("downloaded_image.png")
	if err != nil {
		panic(err)
	} else {
		fmt.Print("Image opened successfully.\n")
	}

	image := gg.NewContextForImage(dc)

	// Get image bounds
	bounds := dc.Bounds()
	// Assign bounds to W and H
	W, H := bounds.Max.X, bounds.Max.Y

	if err := image.LoadFontFace(tempFile.Name(), 72); err != nil {
		panic(err)
	}

	fmt.Printf("Writing text to image.\n")
	image.SetRGB(1, 1, 1)
	image.DrawStringAnchored(s, float64(W/2), float64(H/2), 0.5, 0.5)
	image.Stroke()
	fmt.Printf("Saving file as hello.png\n")
	image.SavePNG("hello.png")
	fmt.Printf("Image text saved successfully.\n")
}
