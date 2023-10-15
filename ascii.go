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
			brightness := (float32(r >> 8)) / 255 // Returns the ratio of brightness
			if brightness >= 0 && brightness < .1 {
				_, err := asciiImg.WriteString(fmt.Sprintf(" "))
				if err != nil {
					panic(err)
				}
			} else if brightness >= .1 && brightness < .2 {
				_, err := asciiImg.WriteString(fmt.Sprintf("."))
				if err != nil {
					panic(err)
				}
			} else if brightness >= .2 && brightness < .3 {
				_, err := asciiImg.WriteString(fmt.Sprintf(":"))
				if err != nil {
					panic(err)
				}
			} else if brightness >= .3 && brightness < .4 {
				_, err := asciiImg.WriteString(fmt.Sprintf(";"))
				if err != nil {
					panic(err)
				}
			} else if brightness >= .4 && brightness < .5 {
				_, err := asciiImg.WriteString(fmt.Sprintf("="))
				if err != nil {
					panic(err)
				}
			} else if brightness >= .5 && brightness < .6 {
				_, err := asciiImg.WriteString(fmt.Sprintf("x"))
				if err != nil {
					panic(err)
				}
			} else if brightness >= .6 && brightness < .7 {
				_, err := asciiImg.WriteString(fmt.Sprintf("X"))
				if err != nil {
					panic(err)
				}
			} else if brightness >= .7 && brightness < .8 {
				_, err := asciiImg.WriteString(fmt.Sprintf("+"))
				if err != nil {
					panic(err)
				}
			} else if brightness >= .8 && brightness < .9 {
				_, err := asciiImg.WriteString(fmt.Sprintf("$"))
				if err != nil {
					panic(err)
				}
			} else {
				_, err := asciiImg.WriteString(fmt.Sprintf("&"))
				if err != nil {
					panic(err)
				}
			}

			_, errors := fo.WriteString(fmt.Sprintf("Pixel at (%d, %d) - R: %d, G: %d, B: %d\n", x, y, r>>8, g>>8, b>>8))
			if errors != nil {
				panic(errors)
			}
		}

		_, err := asciiImg.WriteString(fmt.Sprintf("\n"))
		if err != nil {
			panic(err)
		}
	}
	fmt.Printf("File saved successfully.\n")
}
